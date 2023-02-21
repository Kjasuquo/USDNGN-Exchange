package server

import (
	"context"
	"fmt"
	"github.com/kjasuquo/usdngn-exchange/config"
	"github.com/kjasuquo/usdngn-exchange/internal/controller"
	"github.com/kjasuquo/usdngn-exchange/internal/database/exchange"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start() {

	conf := config.InitDBConfigs()

	exchangeMongoDb, err := exchange.NewExchangeMongoDatabaseAdapter(conf)
	if err != nil {
		log.Fatalf("exchange error: %s\n", err)
	}

	h := &controller.Handler{
		DB:     exchangeMongoDb,
		Config: conf,
	}
	router := SetupRouter(h)

	PORT := fmt.Sprintf(":%s", conf.PORT)
	if PORT == ":" {
		PORT = ":8085"
	}
	srv := &http.Server{
		Addr:    PORT,
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("Server started on %s\n", PORT)
	gracefulShutdown(srv)
}

func gracefulShutdown(srv *http.Server) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
