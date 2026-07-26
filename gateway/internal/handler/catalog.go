package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	catalogv1 "github.com/makoto-developer/strata-sample-platform/proto/catalog/v1"
	"github.com/makoto-developer/strata-sample-platform/gateway/internal/upstream"
)

// CatalogHandler は商品系を CatalogService に中継する。
type CatalogHandler struct {
	clients *upstream.Clients
}

func NewCatalogHandler(c *upstream.Clients) *CatalogHandler { return &CatalogHandler{clients: c} }

func (h *CatalogHandler) Search(c *gin.Context) {
	reply, err := h.clients.Catalog.SearchProducts(c.Request.Context(), &catalogv1.SearchProductsRequest{Id: c.Query("q")})
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reply)
}

func (h *CatalogHandler) Get(c *gin.Context) {
	reply, err := h.clients.Catalog.GetProduct(c.Request.Context(), &catalogv1.GetProductRequest{Id: c.Param("id")})
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reply)
}
