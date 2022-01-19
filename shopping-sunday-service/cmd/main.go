package main

import "shopping-sunday-service/internal/service"

func main() {
	api := service.RestApi{Config: service.Config{
		Port: "8080",
	}}

	api.Start()
}
