package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	orderv1 "github.com/makoto-developer/strata-sample-platform/proto/order/v1"
	"github.com/makoto-developer/strata-sample-platform/gateway/internal/upstream"
)

// OrderHandler は /api/orders 系を OrderService に中継する。
type OrderHandler struct {
	clients *upstream.Clients
}

func NewOrderHandler(c *upstream.Clients) *OrderHandler { return &OrderHandler{clients: c} }

func (h *OrderHandler) Create(c *gin.Context) {
	reply, err := h.clients.Orders.CreateOrder(c.Request.Context(), &orderv1.CreateOrderRequest{})
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, reply)
}

func (h *OrderHandler) Get(c *gin.Context) {
	reply, err := h.clients.Orders.GetOrder(c.Request.Context(), &orderv1.GetOrderRequest{Id: c.Param("id")})
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reply)
}

func (h *OrderHandler) Cancel(c *gin.Context) {
	reply, err := h.clients.Orders.CancelOrder(c.Request.Context(), &orderv1.CancelOrderRequest{Id: c.Param("id")})
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reply)
}

func (h *OrderHandler) ListForUser(c *gin.Context) {
	reply, err := h.clients.Orders.ListOrders(c.Request.Context(), &orderv1.ListOrdersRequest{Id: c.Param("id")})
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reply)
}
