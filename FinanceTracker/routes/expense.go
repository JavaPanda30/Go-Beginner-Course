package routes

import (
	"net/http"
	"strconv"

	"example.com/financetracker/db"
	"example.com/financetracker/models"
	"github.com/gin-gonic/gin"
)

func AddExpense(c *gin.Context) {
	var expense models.Expense
	err := c.ShouldBindBodyWithJSON(&expense)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Mismatch/Parsing Error", "reason": err.Error()})
		return
	}
	inserted, err := expense.Create(db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database Error", "reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Expense Added Successfully", "inserted": &inserted})
}

func GetAllExpense(c *gin.Context) {
	var expense []models.Expense
	expense, err := models.GetExp(db.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"income_record": expense})
}

func GetExpenseById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Mismatch/Parsing Error", "reason": err.Error()})
		return
	}
	expense, err := models.GetIdExp(db.DB, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"income_record": expense})
}

func UpdateExpenseRecord(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Mismatch/Parsing Error", "reason": err.Error()})
		return
	}
	expense, err := models.GetId(db.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense record not found/Does Not Exist", "reason": err.Error()})
		return
	}
	var updated models.Expense
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse given data", "error": err.Error()})
		return
	}
	updated.ID = expense.ID
	_, err = (&updated).PutValues(db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user updated Successfully", "updated_record": updated})
}

func DeleteExpense(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	expense, err := models.GetIdExp(db.DB, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.DeleteExp(db.DB, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record Deleted Successfully", "deleted_record": expense})
}
