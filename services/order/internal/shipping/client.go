package shipping

import (
	"bytes"
	"context"
	"net/http"
)

// Client は shipping サービス(Python / FastAPI)への REST クライアント。
type Client struct {
	base string
}

func New(base string) *Client { return &Client{base: base} }

// CreateShipment は出荷を起票する。
func (c *Client) CreateShipment(ctx context.Context, orderID string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.base+"/shipments", bytes.NewBufferString(orderID))
	if err != nil {
		return err
	}
	_, err = http.DefaultClient.Do(req)
	return err
}

// GetShipment は出荷状況を問い合わせる。
func (c *Client) GetShipment(ctx context.Context, id string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.base+"/shipments/"+id, nil)
	if err != nil {
		return err
	}
	_, err = http.DefaultClient.Do(req)
	return err
}
