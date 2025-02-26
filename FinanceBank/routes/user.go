package routes

import (
	"net/http"

	"example.com/financetracker/db"
	"example.com/financetracker/models"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindBodyWithJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Mismatch/Parsing Error", "reason": err.Error()})
		return
	}
	inserted, err := user.Create(db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database Error", "reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User Added Successfully", "inserted": &inserted})
}

func GetAllUser(c *gin.Context) {
	var user []models.User
	user, err := models.GetUsers(db.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Users_record": user})
}

func GetUserById(c *gin.Context) {
	id := c.Params.ByName("user_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Mismatch/Parsing Error"})
		return
	}
	user, err := models.GetIdUser(db.DB, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user_record": user})
}

func UpdateUserRecord(c *gin.Context) {
	id := c.Params.ByName("user_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Mismatch/Parsing Error"})
		return
	}
	user, err := models.GetIdUser(db.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User record not found/Does Not Exist", "reason": err.Error()})
		return
	}
	var updated models.User
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse given data", "error": err.Error()})
		return
	}
	updated.ID = user.ID
	updated.Account_id = user.Account_id
	_, err = (&updated).PutValues(db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user updated Successfully", "updated_record": updated})
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("user_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parsing error"})
		return
	}
	user, err := models.GetIdUser(db.DB, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.DeleteUser(db.DB, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record Deleted Successfully", "deleted_record": user})
}

func GetTotalIncomeOfUser(c *gin.Context) {
	id := c.Params.ByName("user_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not parse user id"})
		return
	}
	user, err := models.GetIncome(db.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var sum int64
	for _, user := range user {
		if user.UserId == id {
			sum += int64(user.Amount)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Fetch Successfull", "user_id": id, "total_income": sum})
}

func GetTotalExpenseOfUser(c *gin.Context) {
	id := c.Params.ByName("user_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not parse user id"})
	}
	exp, err := models.GetExp(db.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var sum int64
	for _, exp := range exp {
		if exp.UserId == id {

			sum += int64(exp.Amount)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Fetch Successfull", "user_id": id, "total_expense": sum})
}

func GetBalance(c *gin.Context) {
	id := c.Params.ByName("user_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not parse user id"})
		return
	}
	user, err := models.GetIncome(db.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var sum int64
	for _, user := range user {
		if user.UserId == id {
			sum += int64(user.Amount)
		}
	}
	exp, err := models.GetExp(db.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, exp := range exp {
		if exp.UserId == id {
			sum -= int64(exp.Amount)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Fetch Successfull", "user_id": id, "total_balance": sum})
}
