package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/liujunren93/share/client"
	"github.com/liujunren93/share/registry"
	"github.com/liujunren93/share/registry/etcd"
	client2 "github.com/liujunren93/tcc/client"
	"github.com/liujunren93/tcc/test/proto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func main() {
	var err error
	defer func() {
		fmt.Println(err)
	}()
	newClient := client.NewClient()
	newRegistry := etcd.NewRegistry()
	newRegistry.Init(registry.WithAddrs("127.0.0.1:2379"))
	newClient.Init(client.WithRegistry(newRegistry))
	registryServer, err := newClient.Client("tcc.registry")
	registryconn, err := registryServer()
	if err != nil {
		panic(err)
	}
	tccClient, err := client2.NewTccClient(client2.WithGrpcClient(registryconn), client2.WithDB(DB()))
	server1, err := newClient.Client("tcc.server1")
	conn, err := server1()
	server2, err := newClient.Client("tcc.server2")
	conn2, err := server2()
	server1TccClient := proto.NewServer1TccClient(conn)
	server2TccClient := proto.NewServer2TccClient(conn2)

	err = tccClient.Try(
		client2.NewTccOption(server2TccClient, "test.server2", map[string]string{"aaa": "bbb"}, context.TODO()),
		client2.NewTccOption(server1TccClient, "test.server1", map[string]string{"aaa": "bbb"}, context.TODO()),
	)
	fmt.Println(err)
}

func DB() *gorm.DB {
	my := "root:root@tcp(127.0.0.1:3306)/tcc?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       my,    // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	db.Logger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		})

	sqlDB, err := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		panic(err)
	}
	return db.Debug()
}
