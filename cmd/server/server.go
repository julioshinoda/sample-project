package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/julioshinoda/sample-project/internal/api/handler"
	"github.com/julioshinoda/sample-project/internal/api/usecase/customer"
	"github.com/julioshinoda/sample-project/internal/api/usecase/order"
	"github.com/julioshinoda/sample-project/pkg/logger"
	"github.com/julioshinoda/sample-project/pkg/serverconnector"
	"go.uber.org/fx"
)

func main() {

	fx.New(
		fx.Provide(logger.NewZapLogger),
		fx.Provide(serverconnector.NewConnector),
		fx.Provide(serverconnector.NewCustomerConnectorClient),
		fx.Provide(customer.NewService),
		fx.Provide(serverconnector.NewOrdersConnectorClient),
		fx.Provide(order.NewService),
		fx.Provide(handler.NewOrderHandlers),
		fx.Invoke(runHttpServer),
	).Run()

}

func runHttpServer(lifecycle fx.Lifecycle, orderHandler *handler.OrderHandler) {
	lifecycle.Append(fx.Hook{OnStart: func(context.Context) error {
		addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
		fmt.Printf("Starting server on %v\n", addr)
		go http.ListenAndServe(addr, router(orderHandler))
		return nil
	}})
}

func router(orderHandler *handler.OrderHandler) http.Handler {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Route("/orders", func(r chi.Router) {
		handler.MakeOrderHandlers(r, orderHandler)
	})

	return r
}
