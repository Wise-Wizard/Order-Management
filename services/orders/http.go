package main

import (
	"log"
	"net/http"

	handler "github.com/Wise-Wizard/Order-Management/services/orders/handler/orders"
	"github.com/Wise-Wizard/Order-Management/services/orders/service"
)

type HTTPServer struct {
	addr string
}

func NewHTTPServer(addr string) *HTTPServer {
	return &HTTPServer{addr: addr}
}

func (s *HTTPServer) Run() error {
	router := http.NewServeMux()
	orderService := service.NewOrderService()
	orderHandler := handler.NewHTTPOrdersHandler(orderService)
	orderHandler.RegisterRouter(router)
	log.Println("HTTP server is running on port", s.addr)
	return http.ListenAndServe(s.addr, router)
}
