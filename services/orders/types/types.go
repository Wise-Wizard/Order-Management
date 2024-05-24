package types

import (
	"context"

	"github.com/Wise-Wizard/Order-Management/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}
