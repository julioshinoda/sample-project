package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/julioshinoda/sample-project/entity"
	"github.com/julioshinoda/sample-project/internal/api/presenter"
	"github.com/julioshinoda/sample-project/internal/api/usecase/customer"
	"github.com/julioshinoda/sample-project/internal/api/usecase/order"
	"go.uber.org/zap"
)

type OrderHandler struct {
	logger      *zap.Logger
	customerSrv customer.UseCase
	orderSrv    order.UseCase
}

func NewOrderHandlers(logger *zap.Logger, customerSrv customer.UseCase, orderSrv order.UseCase) *OrderHandler {
	return &OrderHandler{
		logger:      logger,
		customerSrv: customerSrv,
		orderSrv:    orderSrv,
	}
}

func MakeOrderHandlers(r chi.Router, c *OrderHandler) {
	r.Post("/", c.FindOrders)

}

func (c *OrderHandler) FindOrders(w http.ResponseWriter, r *http.Request) {
	var email entity.CustomerEmail
	if err := json.NewDecoder(r.Body).Decode(&email); err != nil {
		c.logger.Error(
			"Could not find a customer",
			zap.Error(err),
		)
		response, _ := json.Marshal(presenter.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		presenter.MountResponse(w, http.StatusBadRequest, response)
		return
	}
	customer, err := c.customerSrv.GetByEmail(email.Email)
	if err != nil || len(customer) == 0 {
		c.logger.Error(
			"Could not find a customer",
			zap.String("email", email.Email),
			zap.Error(err),
		)
		response, _ := json.Marshal(presenter.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		presenter.MountResponse(w, http.StatusBadRequest, response)
		return
	}

	orderResp, err := c.orderSrv.GetByCustomerID(customer[0].Customerid)
	if err != nil {
		c.logger.Error(
			"Could not find orders",
			zap.String("email", email.Email),
			zap.Error(err),
		)
		response, _ := json.Marshal(presenter.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		presenter.MountResponse(w, http.StatusBadRequest, response)
		return
	}
	response, _ := json.Marshal(orderResp)
	presenter.MountResponse(w, http.StatusOK, response)
}
