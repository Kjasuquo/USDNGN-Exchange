package main

import (
	"creditsystem/cmd/server"
	_ "creditsystem/docs"
	"fmt"
	"log"
)

// @title Credit-System API
// @version 1.0
// @description This is a Propchain API endpoints for Users.
// @contact.name Propchain Admin
// @contact.email support@propchain.com
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name API-KEY
// @securityDefinitions.apikey ApiSecretAuth
// @in header
// @name API-SECRET
func main() {
	fmt.Println("Hello, world!")

	server.Start()
	//go controller.CheckTheStatusOfPendingTransaction()
	//go exchange.SendExchangeRequestCronJob()
	log.Println("success")
}
