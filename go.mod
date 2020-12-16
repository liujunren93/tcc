module github.com/liujunren93/tcc

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.3
	github.com/liujunren93/share v0.0.0-20200922022710-0eb46ed2fd91
	golang.org/x/net v0.0.0-20201016165138-7b1cca2348c0 // indirect
	google.golang.org/grpc v1.33.0
	gorm.io/driver/mysql v1.0.2
	gorm.io/gorm v1.20.2
)
