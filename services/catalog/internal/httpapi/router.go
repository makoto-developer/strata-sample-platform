package httpapi

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/makoto-developer/strata-sample-platform/services/catalog/internal/handler"
)

// Serve は商品参照の公開 REST(echo)。
func Serve(addr string, srv *handler.Server) {
	e := echo.New()
	api := e.Group("/api/v1")
	api.GET("/products", listProducts)
	api.GET("/products/:id", getProduct)
	api.POST("/products", upsertProduct)
	_ = e.Start(addr)
}

func listProducts(c echo.Context) error { return c.JSON(http.StatusOK, []string{}) }

func getProduct(c echo.Context) error { return c.JSON(http.StatusOK, map[string]string{"id": c.Param("id")}) }

func upsertProduct(c echo.Context) error { return c.NoContent(http.StatusAccepted) }
