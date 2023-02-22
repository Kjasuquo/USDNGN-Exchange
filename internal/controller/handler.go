package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kjasuquo/usdngn-exchange/config"
	"github.com/kjasuquo/usdngn-exchange/internal/models"
	"github.com/kjasuquo/usdngn-exchange/internal/port"
)

type Handler struct {
	DB     port.DB
	Rates  port.Rates
	Config config.Config
}

func (h *Handler) Ping(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (h *Handler) GetUserFromContext(c *gin.Context) (*models.User, error) {
	authUser, exists := c.Get("user")
	if !exists {
		return nil, fmt.Errorf("error getting user from context")
	}
	user, ok := authUser.(*models.User)
	if !ok {
		return nil, fmt.Errorf("an error occurred")
	}
	return user, nil
}
