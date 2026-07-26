package main

import (
	"log"

	"github.com/makoto-developer/strata-sample-platform/services/order/internal/client"
	"github.com/makoto-developer/strata-sample-platform/services/order/internal/config"
	"github.com/makoto-developer/strata-sample-platform/services/order/internal/handler"
	"github.com/makoto-developer/strata-sample-platform/services/order/internal/usecase"
)

func main() {
	cfg := config.Load()
	clients := client.Dial(cfg)
	uc := usecase.New(clients)
	log.Fatal(handler.New(uc).Serve(cfg.GRPCAddr))
}
