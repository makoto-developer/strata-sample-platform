package config

import "os"

type Config struct {
	GRPCAddr    string
	HTTPAddr    string
	DatabaseURL string
	AuthSecret  string
}

func Load() Config {
	return Config{
		GRPCAddr:    env("USER_GRPC_ADDR", ":50051"),
		HTTPAddr:    env("USER_HTTP_ADDR", ":8081"),
		DatabaseURL: env("USER_DATABASE_URL", "postgres://localhost/users"),
		AuthSecret:  env("AUTH_SECRET", "dev"),
	}
}

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
