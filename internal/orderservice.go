package internal

import (
	"context"
	"log"

	"github.com/koyeb/go-grpc/protogen/golang/orders"
)

// OrderService should implement the OrdersServer interface generated from grpc.
//
// UnimplementedOrdersServer must be embedded to have forwarded compatible implementations.
type OrderService struct {
	db *DB
	orders.UnimplementedOrdersServer
}

// NewOrderService creates a new OrderService
func NewOrderService(db *DB) OrderService {
	return OrderService{db: db}
}

// AddOrder implements the AddOrder method of the grpc OrdersServer interface to add a new order
func (o *OrderService) AddOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.Empty, error) {
	log.Printf("Received an add-order request")

	err := o.db.AddOrder(req.GetOrder())

	return &orders.Empty{}, err
}
