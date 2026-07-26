package client

import (
	inventoryv1 "github.com/makoto-developer/strata-sample-platform/proto/inventory/v1"
	notificationv1 "github.com/makoto-developer/strata-sample-platform/proto/notification/v1"
	paymentv1 "github.com/makoto-developer/strata-sample-platform/proto/payment/v1"
	"github.com/makoto-developer/strata-sample-platform/services/order/internal/config"
)

// Clients は order が依存する下流サービスの gRPC クライアント束。
type Clients struct {
	Payment      paymentv1.PaymentServiceClient
	Inventory    inventoryv1.InventoryServiceClient
	Notification notificationv1.NotificationServiceClient
	ShippingURL  string
}

func Dial(cfg config.Config) *Clients {
	return &Clients{ShippingURL: cfg.ShippingURL}
}
