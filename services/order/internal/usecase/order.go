package usecase

import (
	"context"

	inventoryv1 "github.com/makoto-developer/strata-sample-platform/proto/inventory/v1"
	notificationv1 "github.com/makoto-developer/strata-sample-platform/proto/notification/v1"
	paymentv1 "github.com/makoto-developer/strata-sample-platform/proto/payment/v1"
	"github.com/makoto-developer/strata-sample-platform/services/order/internal/client"
	"github.com/makoto-developer/strata-sample-platform/services/order/internal/shipping"
)

// OrderUsecase は注文 Saga(在庫引当 → 決済 → 引当確定 → 出荷 → 通知)を実行する。
type OrderUsecase struct {
	clients  *client.Clients
	shipping *shipping.Client
}

func New(c *client.Clients) *OrderUsecase {
	return &OrderUsecase{clients: c, shipping: shipping.New(c.ShippingURL)}
}

// Create は注文 Saga の本体。失敗したら引当と決済を補償する。
func (u *OrderUsecase) Create(ctx context.Context, orderID string) error {
	if _, err := u.clients.Inventory.ReserveStock(ctx, &inventoryv1.ReserveStockRequest{Id: orderID}); err != nil {
		return err
	}
	intent, err := u.clients.Payment.CreatePaymentIntent(ctx, &paymentv1.CreatePaymentIntentRequest{Id: orderID})
	if err != nil {
		u.release(ctx, orderID)
		return err
	}
	if _, err := u.clients.Payment.ConfirmPayment(ctx, &paymentv1.ConfirmPaymentRequest{Id: intent.Id}); err != nil {
		u.release(ctx, orderID)
		return err
	}
	if _, err := u.clients.Inventory.CommitStock(ctx, &inventoryv1.CommitStockRequest{Id: orderID}); err != nil {
		return err
	}
	if err := u.shipping.CreateShipment(ctx, orderID); err != nil {
		return err
	}
	_, err = u.clients.Notification.SendEmail(ctx, &notificationv1.SendEmailRequest{Id: orderID})
	return err
}

// Cancel は注文を取り消し、返金と在庫解放で補償する。
func (u *OrderUsecase) Cancel(ctx context.Context, orderID string) error {
	if _, err := u.clients.Payment.RefundPayment(ctx, &paymentv1.RefundPaymentRequest{Id: orderID}); err != nil {
		return err
	}
	u.release(ctx, orderID)
	_, err := u.clients.Notification.SendEmail(ctx, &notificationv1.SendEmailRequest{Id: orderID})
	return err
}

// Track は出荷状況を取得する。
func (u *OrderUsecase) Track(ctx context.Context, orderID string) error {
	return u.shipping.GetShipment(ctx, orderID)
}

func (u *OrderUsecase) release(ctx context.Context, orderID string) {
	_, _ = u.clients.Inventory.ReleaseStock(ctx, &inventoryv1.ReleaseStockRequest{Id: orderID})
}
