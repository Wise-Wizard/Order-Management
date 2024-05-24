package handler

import (
	"context"

	orders "github.com/Wise-Wizard/Order-Management/services/common/genproto/orders"
	"github.com/Wise-Wizard/Order-Management/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGRPCHandler struct {
	orderService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGRPCOrdersService(gRPCServer *grpc.Server, orderService types.OrderService) {
	gRPCHandler := &OrdersGRPCHandler{
		orderService: orderService,
	}
	orders.RegisterOrderServiceServer(gRPCServer, gRPCHandler)
}

func (h *OrdersGRPCHandler) GetOrders(ctx context.Context, req *orders.GetOrdersRequest) (*orders.GetOrdersResponse, error) {
	orders, err := h.orderService.GetOrders(ctx, req)
	if err != nil {
		return nil, err
	}
	res := &orders.GetOrdersResponse{
		Orders: orders,
	}
	return res, nil
}

func (h *OrdersGRPCHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    42,
		CustomerID: 2,
		ProductID:  10,
		Quantity:   1,
	}
	err := h.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	res := &orders.CreateOrderResponse{
		Status: "Order Created Succesfully",
	}
	return res, nil
}
