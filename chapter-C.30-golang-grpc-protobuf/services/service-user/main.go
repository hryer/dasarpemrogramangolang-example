package main

import (
	"chapter-c30/common/config"
	"chapter-c30/common/model"
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var localStorage *model.UserList

func init() {
	localStorage = new(model.UserList)
	localStorage.List = make([]*model.User, 0)
}

//UsersServer is ...
type UsersServer struct {
	model.UnimplementedUsersServer
}

//Register is ...
func (UsersServer) Register(ctx context.Context, param *model.User) (*emptypb.Empty, error) {
	localStorage.List = append(localStorage.List, param)
	log.Println("Registering user", param.String())

	return new(empty.Empty), nil
}

//List is ...
func (UsersServer) List(ctx context.Context, void *empty.Empty) (*model.UserList, error) {
	return localStorage, nil
}

func main() {
	srv := grpc.NewServer()
	var usrSrv UsersServer
	model.RegisterUsersServer(srv, usrSrv)

	log.Println("Starting RPC server at", config.SERVICE_USER_PORT)

	l, err := net.Listen("tcp", config.SERVICE_USER_PORT)
	if err != nil {
		log.Fatalf("Could not listen to %s: %v", config.SERVICE_USER_PORT, err.Error())
	}
	log.Fatal(srv.Serve(l))
}
