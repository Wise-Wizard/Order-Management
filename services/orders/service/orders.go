package service

import (
	"context"

	orders "github.com/Wise-Wizard/Order-Management/services/common/genproto/orders"
)

var ordersDB = make([]*orders.Order, 0)

type OrderService struct {
	// Store
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersDB = append(ordersDB, order)
	return nil
}

func (s *OrderService) GetOrders(ctx context.Context, req *orders.GetOrdersRequest) ([]*orders.Order, error) {
	return ordersDB, nil
}