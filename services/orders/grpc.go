package main

import (
	handler "grpc-go/services/orders/handler/orders"
	"grpc-go/services/orders/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

type grpcServer struct {
	addr string
}

func NewGRPCServer(addr string) *grpcServer {
	return &grpcServer{addr: addr}
}

func (s *grpcServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	// register our grpc services
	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(grpcServer, orderService)

	log.Println("Starting on gRPC server on", s.addr)
	return grpcServer.Serve(lis)
}
