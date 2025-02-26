package routes

import (
	"example.com/eventbook/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//Event Route
	server.GET("/events", GetEvents)
	server.GET(`/events/:id`, GetEventsById)
	server.POST("/events", CreateEvent)
	server.PUT(`/events/:id`, UpdateEvent)
	server.DELETE(`/events/:id`, DeleteEvent)

	//User Routes
	server.GET("/users", GetUsers)
	// server.GET("/user/event/:id", GetUserEvents)
	server.POST("/signup", CreateUser)
	server.POST("/login", Login)
	server.DELETE("/user/:id", DeleteUser)

	//Protected Route
	protected := server.Group("/")
	protected.Use(middleware.AuthToken)
	protected.POST("/events", CreateEvent)
	protected.PUT(`/events/:id`, UpdateEvent)
	protected.DELETE(`/events/:id`, DeleteEvent)

	//Registrations
	protected.POST("/events/:id/register", AddRegistration)
	protected.DELETE("/events/:id/register", CancelRegistration)
}
