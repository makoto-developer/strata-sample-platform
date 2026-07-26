package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	userv1 "github.com/makoto-developer/strata-sample-platform/proto/user/v1"
	"github.com/makoto-developer/strata-sample-platform/gateway/internal/upstream"
)

// UserHandler は /api/users 系を gRPC の UserService に中継する。
type UserHandler struct {
	clients *upstream.Clients
}

func NewUserHandler(c *upstream.Clients) *UserHandler { return &UserHandler{clients: c} }

func (h *UserHandler) Get(c *gin.Context) {
	reply, err := h.clients.Users.GetUser(c.Request.Context(), &userv1.GetUserRequest{Id: c.Param("id")})
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reply)
}

func (h *UserHandler) List(c *gin.Context) {
	reply, err := h.clients.Users.ListUsers(c.Request.Context(), &userv1.ListUsersRequest{})
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reply)
}

func (h *UserHandler) Create(c *gin.Context) {
	reply, err := h.clients.Users.CreateUser(c.Request.Context(), &userv1.CreateUserRequest{})
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, reply)
}
