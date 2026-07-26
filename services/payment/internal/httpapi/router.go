package httpapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/makoto-developer/strata-sample-platform/services/payment/internal/handler"
)

// Routes は決済プロバイダからの webhook 受信口。
func Routes(srv *handler.Server) http.Handler {
	r := chi.NewRouter()
	r.Post("/webhooks/payment-provider", handleProviderWebhook)
	r.Get("/health", handleHealth)
	return r
}

// handleProviderWebhook は課金結果の非同期通知(外部 → 自システム)を受ける。
func handleProviderWebhook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("ok"))
}
