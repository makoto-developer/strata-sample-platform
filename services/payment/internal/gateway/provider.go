package gateway

import (
	"bytes"
	"context"
	"net/http"
	"os"
)

// Provider は外部決済プロバイダ(SaaS)への HTTP クライアント。
type Provider struct {
	client *http.Client
}

func New() *Provider { return &Provider{client: http.DefaultClient} }

// Charge は外部 API に課金を依頼する(自システム → 外部)。
func (p *Provider) Charge(ctx context.Context, amount string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		"https://api.payments.example.com/v1/charges", bytes.NewBufferString(amount))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("PAYMENT_PROVIDER_KEY"))
	_, err = p.client.Do(req)
	return err
}

// Refund は外部 API に返金を依頼する。
func (p *Provider) Refund(ctx context.Context, chargeID string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		"https://api.payments.example.com/v1/refunds", bytes.NewBufferString(chargeID))
	if err != nil {
		return err
	}
	_, err = p.client.Do(req)
	return err
}
