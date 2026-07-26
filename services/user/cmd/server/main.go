package main

import (
	"log"
	"net"
	"net/http"

	"github.com/makoto-developer/strata-sample-platform/services/user/internal/config"
	"github.com/makoto-developer/strata-sample-platform/services/user/internal/httpapi"
	"github.com/makoto-developer/strata-sample-platform/services/user/internal/handler"
	"github.com/makoto-developer/strata-sample-platform/services/user/internal/store"
)

func main() {
	cfg := config.Load()
	db := store.New(cfg.DatabaseURL)
	srv := handler.New(db)

	go func() {
		log.Printf("user http listening on %s", cfg.HTTPAddr)
		if err := http.ListenAndServe(cfg.HTTPAddr, httpapi.Routes(db)); err != nil {
			log.Fatal(err)
		}
	}()

	lis, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("user grpc listening on %s", cfg.GRPCAddr)
	if err := srv.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
