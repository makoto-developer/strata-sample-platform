package main

import (
	"log"
	"net/http"

	"github.com/makoto-developer/strata-sample-platform/graphql/subgraphs/catalog/resolver"
)

func main() {
	r := &resolver.Resolver{}
	_ = r
	http.HandleFunc("POST /graphql", handleGraphQL)
	log.Fatal(http.ListenAndServe(":4003", nil))
}

func handleGraphQL(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
