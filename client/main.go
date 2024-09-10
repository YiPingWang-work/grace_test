package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	service_pb "grpc_test/service.pb"
	"log"
	"os"
	"time"
)

func main() {
	log.Println("pid is", os.Getpid())
	conn, err := grpc.Dial("localhost:9999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	cli := service_pb.NewPersonServiceClient(conn)
	id := 0
	for {
		fmt.Println(".")
		id++
		resp, err := cli.GetInfo(context.Background(), &service_pb.Req{
			Name: fmt.Sprintf("name_%d", id),
		})
		if err != nil {
			log.Println(err, conn.GetState().String(), id)
		} else {
			log.Println(resp, conn.GetState().String(), id)
		}
		time.Sleep(time.Second)
	}
}
