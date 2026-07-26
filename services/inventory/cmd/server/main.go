package main

import (
	"log"

	"github.com/makoto-developer/strata-sample-platform/services/inventory/internal/handler"
)

func main() {
	log.Fatal(handler.New().Serve(":50053"))
}
