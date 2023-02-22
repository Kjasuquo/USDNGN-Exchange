package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kjasuquo/usdngn-exchange/internal/controller"
	"github.com/kjasuquo/usdngn-exchange/internal/middleware"
	"log"
)

func DefineRoutes(handler *controller.Handler) *gin.Engine {
	log.Println("Routes defined")

	router := gin.Default()

	r := router.Group("/api/v1")
	{
		r.POST("/signup", handler.SignUp)
		r.POST("/login", handler.Login)
	}

	authorized := r.Use(middleware.AuthorizeUser(handler))
	{
		authorized.POST("user/profile", handler.UserProfile)
	}

	return router
}

func SetupRouter(h *controller.Handler) *gin.Engine {
	log.Println("Router setup")
	r := DefineRoutes(h)

	return r
}
