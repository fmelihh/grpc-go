package types

import (
	"context"
	"grpc-go/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *orders.Order) error
	GetOrders(ctx context.Context) []*orders.Order
}
