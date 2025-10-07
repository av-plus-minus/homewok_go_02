package main

import (
	"fmt"
	"gateway/internal/controller"
	"gateway/internal/server"
	"log"
)

func main() {
	pingController := controller.NewPingController()

	srv := server.NewServer(pingController)

	fmt.Println("Gateway service starting on :8080")
	log.Fatal(srv.Start(":8080"))
}
