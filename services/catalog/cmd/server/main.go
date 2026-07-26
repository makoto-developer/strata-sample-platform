package main

import (
	"log"

	"github.com/makoto-developer/strata-sample-platform/services/catalog/internal/config"
	"github.com/makoto-developer/strata-sample-platform/services/catalog/internal/handler"
	"github.com/makoto-developer/strata-sample-platform/services/catalog/internal/httpapi"
)

func main() {
	cfg := config.Load()
	srv := handler.New(cfg)
	go httpapi.Serve(cfg.HTTPAddr, srv)
	log.Fatal(srv.Serve(cfg.GRPCAddr))
}
