package mail

import (
	"bytes"
	"context"
	"net/http"
	"os"
)

// Sender は外部メール配信 SaaS への HTTP クライアント。
type Sender struct{}

func New() *Sender { return &Sender{} }

// Send はメール配信 API を叩く(自システム → 外部)。
func (s *Sender) Send(ctx context.Context, to string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		"https://api.mail.example.com/v3/send", bytes.NewBufferString(to))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("MAIL_API_KEY"))
	_, err = http.DefaultClient.Do(req)
	return err
}

// NotifyOps は運用チャンネルへ webhook を送る。
func (s *Sender) NotifyOps(ctx context.Context, text string) error {
	_, err := http.Post("https://hooks.chat.example.com/services/ops", "application/json", bytes.NewBufferString(text))
	return err
}
