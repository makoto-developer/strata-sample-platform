package main

import (
	"log"

	"github.com/makoto-developer/strata-sample-platform/gateway/internal/router"
)

func main() {
	log.Fatal(router.Run(":8080"))
}
