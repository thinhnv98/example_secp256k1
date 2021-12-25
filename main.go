package main

import (
	"example_secp256k1/route"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	//Routes
	router := route.Route{
		Server: server,
	}

	router.Register()

	//Start server
	server.Run()
}
