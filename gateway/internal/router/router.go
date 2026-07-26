package router

import (
	"github.com/gin-gonic/gin"

	"github.com/makoto-developer/strata-sample-platform/gateway/internal/handler"
	"github.com/makoto-developer/strata-sample-platform/gateway/internal/upstream"
)

// Run は BFF の REST ルーティングを定義して待ち受ける。
func Run(addr string) error {
	clients := upstream.Dial()
	users := handler.NewUserHandler(clients)
	orders := handler.NewOrderHandler(clients)
	catalog := handler.NewCatalogHandler(clients)

	r := gin.New()
	api := r.Group("/api")
	api.GET("/users/:id", users.Get)
	api.GET("/users", users.List)
	api.POST("/users", users.Create)
	api.GET("/users/:id/orders", orders.ListForUser)
	api.POST("/orders", orders.Create)
	api.GET("/orders/:id", orders.Get)
	api.POST("/orders/:id/cancel", orders.Cancel)
	api.GET("/products", catalog.Search)
	api.GET("/products/:id", catalog.Get)
	return r.Run(addr)
}
