package service

import (
	"context"
	service_pb2 "grpc_test/service.pb"
	"log"
)

type Service struct {
	service_pb2.UnimplementedPersonServiceServer
}

func (s Service) GetInfo(ctx context.Context, req *service_pb2.Req) (*service_pb2.Resp, error) {
	log.Println("GetInfo", req.String())
	return &service_pb2.Resp{
		Status: true,
	}, nil
}

var _ service_pb2.PersonServiceServer = (*Service)(nil)
