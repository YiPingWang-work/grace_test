package main

import (
	"google.golang.org/grpc"
	"grpc_test/service"
	"grpc_test/service.pb"
	"log"
	"net"
	"os"
)

func main() {
	log.Println("pid is", os.Getpid())
	listen, _ := net.Listen("tcp", ":9999")
	grpcServer := grpc.NewServer()
	service_pb.RegisterPersonServiceServer(grpcServer, &service.Service{})
	log.Println("begin grpc server")
	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}
