package server

import (
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type options struct {
	DB         *gorm.DB
	server     *grpc.Server
	ClientName string
}

func WithDB(db *gorm.DB) option {
	return func(o *options) {
		o.DB = db
	}
}

func WithServer(server *grpc.Server) option {
	return func(o *options) {
		o.server = server
	}
}

func WithClientName(ClientName string) option {
	return func(o *options) {
		o.ClientName = ClientName
	}
}