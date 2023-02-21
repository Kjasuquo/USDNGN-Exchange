package main

import (
	"fmt"
	"github.com/kjasuquo/usdngn-exchange/cmd/server"
	"log"
)

func main() {
	fmt.Println("Hello, world!")

	server.Start()
	log.Println("successfully disconnected")
}
