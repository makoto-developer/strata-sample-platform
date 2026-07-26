package main

import (
	"log"
	"net/http"

	"github.com/makoto-developer/strata-sample-platform/services/payment/internal/handler"
	"github.com/makoto-developer/strata-sample-platform/services/payment/internal/httpapi"
)

func main() {
	srv := handler.New()
	go func() { log.Fatal(http.ListenAndServe(":8083", httpapi.Routes(srv))) }()
	log.Fatal(srv.Serve(":50054"))
}
