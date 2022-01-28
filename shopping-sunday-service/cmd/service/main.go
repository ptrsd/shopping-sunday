package main

import (
	"os"
	"shopping-sunday-service/internal/service"
)

func main() {
	var port = "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	api := service.RestApi{Config: service.Config{
		Port: port,
	}}
	api.Start()
}
