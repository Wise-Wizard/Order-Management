package handler

import (
	"net/http"

	orders "github.com/Wise-Wizard/Order-Management/services/common/genproto/orders"
	"github.com/Wise-Wizard/Order-Management/services/common/util"
	"github.com/Wise-Wizard/Order-Management/services/orders/types"
)

type OrdersHTTPHandler struct {
	orderService types.OrderService
}

func NewHTTPOrdersHandler(orderService types.OrderService) *OrdersHTTPHandler {
	return &OrdersHTTPHandler{
		orderService: orderService,
	}
}

func (h *OrdersHTTPHandler) RegisterRouter(r *http.ServeMux) {
	r.HandleFunc("/orders", h.CreateOrder)
}

func (h *OrdersHTTPHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		util.WriteError(w, http.StatusMethodNotAllowed, http.ErrNotSupported)
		return
	}
	var req orders.CreateOrderRequest
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order{
		OrderID:    42,
		CustomerID: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}

	err = h.orderService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{Status: "success"}
	util.WriteJSON(w, http.StatusOK, res)
}
