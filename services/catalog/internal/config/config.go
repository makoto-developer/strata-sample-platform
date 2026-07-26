package config

import "os"

type Config struct {
	GRPCAddr  string
	HTTPAddr  string
	SearchURL string
}

func Load() Config {
	return Config{
		GRPCAddr:  env("CATALOG_GRPC_ADDR", ":50052"),
		HTTPAddr:  env("CATALOG_HTTP_ADDR", ":8082"),
		SearchURL: env("SEARCH_SERVICE_URL", "http://search:8087"),
	}
}

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
