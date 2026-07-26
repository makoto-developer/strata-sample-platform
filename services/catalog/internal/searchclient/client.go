package searchclient

import (
	"bytes"
	"context"
	"net/http"
)

// Client は search サービス(REST)の HTTP クライアント。
type Client struct {
	base string
	http *http.Client
}

func New(base string) *Client {
	return &Client{base: base, http: http.DefaultClient}
}

// Query は全文検索を投げる。
func (c *Client) Query(ctx context.Context, q string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.base+"/search/products?q="+q, nil)
	if err != nil {
		return err
	}
	_, err = c.http.Do(req)
	return err
}

// Index は検索インデックスを更新する。
func (c *Client) Index(ctx context.Context, id string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.base+"/search/index", bytes.NewBufferString(id))
	if err != nil {
		return err
	}
	_, err = c.http.Do(req)
	return err
}
