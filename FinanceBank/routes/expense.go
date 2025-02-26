package routes

import (
	"net/http"

	"example.com/financetracker/db"
	"example.com/financetracker/models"
	"github.com/gin-gonic/gin"
)

func AddExpense(c *gin.Context) {
	var expense models.Expense
	err := c.ShouldBindJSON(&expense)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Mismatch/Parsing Error", "reason": err.Error()})
		return
	}
	inserted, err := expense.Create(db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database Error", "reason": err.Error()})
		return
	}
	amt, err := models.GetAccountAmount(inserted.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database Error", "reason": err.Error()})
		return
	}
	amt = amt - inserted.Amount
	models.UpdateAccountAmt(amt, expense.UserId)
	c.JSON(http.StatusOK, gin.H{"message": "Expense Added Successfully", "inserted": &inserted})
}

func GetAllExpense(c *gin.Context) {
	var expense []models.Expense
	expense, err := models.GetExp(db.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"expense_record": expense})
}

func GetExpenseById(c *gin.Context) {
	id := c.Params.ByName("expense_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Mismatch/Parsing Error", "reason": "parsing error"})
		return
	}
	expense, err := models.GetIdExp(db.DB, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"expense_record": expense})
}

func UpdateExpenseRecord(c *gin.Context) {
	id := c.Params.ByName("expense_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Mismatch/Parsing Error", "reason": "parsing error"})
		return
	}
	expense, err := models.GetIdExp(db.DB, id)
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
	updated.UserId = expense.UserId
	_, err = (&updated).PutValues(db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Expense record updated Successfully", "updated_record": updated})
}

func DeleteExpense(c *gin.Context) {
	id := c.Params.ByName("expense_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parsing error"})
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
