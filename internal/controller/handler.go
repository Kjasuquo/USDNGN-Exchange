package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kjasuquo/usdngn-exchange/config"
	"github.com/kjasuquo/usdngn-exchange/internal/port"
)

type Handler struct {
	DB     port.DB
	Config config.Config
}

func (h *Handler) Ping(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
