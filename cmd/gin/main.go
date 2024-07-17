package main

import (
	"log"

	"github.com/codescalersinternships/datetime-server-salmaelsoly/pkg/gin"
)

func main() {
	server := routes.SetUpRouter()
	err := server.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
