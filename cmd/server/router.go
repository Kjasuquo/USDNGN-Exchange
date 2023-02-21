package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kjasuquo/usdngn-exchange/internal/controller"
	"log"
)

func DefineRoutes(handler *controller.Handler) *gin.Engine {
	log.Println("Routes defined")

	router := gin.Default()

	//r := router.Group("/api/v1")
	{

	}

	return router
}

func SetupRouter(h *controller.Handler) *gin.Engine {
	log.Println("Router setup")
	r := DefineRoutes(h)

	return r
}
