package routes

import (
	"net/http"
	"strconv"

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
	c.JSON(http.StatusOK, gin.H{"message": "Income Added Successfully", "inserted": &inserted})
}

func GetAllIncome(c *gin.Context) {
	var income []models.Income
	income, err := models.Get(db.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"income_record": income})
}

func GetIncomeById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Mismatch/Parsing Error", "reason": err.Error()})
		return
	}
	income, err := models.GetId(db.DB, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"income_record": income})
}

func UpdateIncomeRecord(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input Mismatch/Parsing Error", "reason": err.Error()})
		return
	}
	income, err := models.GetId(db.DB, id)
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
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	income, err := models.GetId(db.DB, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.Delete(db.DB, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record Deleted Successfully", "deleted_record": income})
}

func GetBalance(){

}