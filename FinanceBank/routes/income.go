package routes

import (
	"net/http"

	"example.com/financetracker/db"
	"example.com/financetracker/models"
	"github.com/gin-gonic/gin"
)

func AddIncome(c *gin.Context) {
	var income models.Income
	err := c.ShouldBindBodyWithJSON(&income)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Mismatch/Parsing Error", "reason": err.Error()})
		return
	}
	inserted, err := income.Create(db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database Error", "reason": err.Error()})
		return
	}
	amt, err := models.GetAccountAmount(inserted.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database Error", "reason": err.Error()})
		return
	}
	amt = inserted.Amount + amt
	models.UpdateAccountAmt(amt, income.UserId)
	c.JSON(http.StatusOK, gin.H{"message": "Income Added Successfully", "inserted": &inserted})
}

func UpdateIncomeRecord(c *gin.Context) {
	id := c.Params.ByName("income_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Mismatch/Parsing Error"})
		return
	}
	income, err := models.GetIncomeById(db.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Income record not found/Does Not Exist", "reason": err.Error()})
		return
	}
	var updated models.Income
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse given data", "error": err.Error()})
		return
	}
	updated.ID = income.ID
	_, err = (&updated).PutValues(db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user updated Successfully", "updated_record": updated})
}

func DeleteIncome(c *gin.Context) {
	id := c.Params.ByName("income_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parsing Error"})
		return
	}
	income, err := models.GetIncomeById(db.DB, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.DeleteIncome(db.DB, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record Deleted Successfully", "deleted_record": income})
}

func GetAllIncome(c *gin.Context) {
	var income []models.Income
	income, err := models.GetIncome(db.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"income_record": income})
}

func GetIncomeById(c *gin.Context) {
	id := c.Params.ByName("income_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Mismatch/Parsing Error"})
		return
	}
	income, err := models.GetIncomeById(db.DB, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"income_record": income})
}
