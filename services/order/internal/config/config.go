package config

import "os"

type Config struct {
	GRPCAddr        string
	PaymentAddr     string
	InventoryAddr   string
	NotificationAddr string
	ShippingURL     string
}

func Load() Config {
	return Config{
		GRPCAddr:         env("ORDER_GRPC_ADDR", ":50055"),
		PaymentAddr:      env("PAYMENT_SERVICE_ADDR", "payment:50054"),
		InventoryAddr:    env("INVENTORY_SERVICE_ADDR", "inventory:50053"),
		NotificationAddr: env("NOTIFICATION_SERVICE_ADDR", "notification:50056"),
		ShippingURL:      env("SHIPPING_SERVICE_URL", "http://shipping:8086"),
	}
}

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
