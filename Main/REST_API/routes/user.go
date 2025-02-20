package routes

import (
	"net/http"
	"strconv"

	util "example.com/eventbook/Util"
	"example.com/eventbook/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse given data", "error": err.Error()})
		return
	}
	user.ID = 1
	if err := user.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save User", "error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User Created Successfully", "user": user})
}

func GetUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "User Cannot be Fetched", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully Fetched Users",
		"Users":   users,
	})
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}
	user, err := models.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot Delete User", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User Deleted Successfully", "user": user})
}

func Login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User Credential cannot be parsed", "error": err.Error()})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Credentials", "error": err.Error()})
		return
	}
	token, err := util.GetToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Credentials", "error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "Login Successfull", "user": user, "token": token})
}
