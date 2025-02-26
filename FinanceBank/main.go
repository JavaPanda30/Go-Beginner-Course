package main

import (
	"example.com/financetracker/db"
	"example.com/financetracker/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.Routes(server)
	server.Run()
}
