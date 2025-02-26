package middleware

import (
	"net/http"

	util "example.com/eventbook/Util"
	"github.com/gin-gonic/gin"
)

func AuthToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Login to Use This Route "})
	}
	id, err := util.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Login Again to Use This Route", "error": err.Error()})
	}
	c.Set("id", id)
	c.Next()
}
