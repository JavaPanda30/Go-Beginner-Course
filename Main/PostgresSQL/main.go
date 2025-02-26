package main

import (
	"example.com/eventbook/db"
	"example.com/eventbook/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default() //starts an engine - starts a http server
	routes.RegisterRoutes(server)
	server.Run()
}
