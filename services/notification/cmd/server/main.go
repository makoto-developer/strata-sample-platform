package main

import (
	"log"

	"github.com/makoto-developer/strata-sample-platform/services/notification/internal/handler"
)

func main() {
	log.Fatal(handler.New().Serve(":50056"))
}
