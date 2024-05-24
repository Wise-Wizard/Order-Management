package main

import (
	"log"
	"net"

	handler "github.com/Wise-Wizard/Order-Management/services/orders/handler/orders"
	"github.com/Wise-Wizard/Order-Management/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServcer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	gRPCServer := grpc.NewServer()
	log.Println("gRPC server is running on port", s.addr)

	// Register Our Services
	orderService := service.NewOrderService()
	handler.NewGRPCOrdersService(gRPCServer, orderService)
	return gRPCServer.Serve(lis)
}
