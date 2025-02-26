package routes

import (
	"net/http"

	"example.com/financetracker/db"
	"example.com/financetracker/models"
	"github.com/gin-gonic/gin"
)

func SendMoney(c *gin.Context) {
	var tr models.Transaction
	if err := c.ShouldBindBodyWithJSON(&tr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"parsing error": err.Error()})
		return
	}
	added, err := tr.Insert(db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"insert DB error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "true", "Transaction": added})
}

func GetReceipt(c *gin.Context) {
	id := c.Param("transaction_id")
	receipt, err := models.GetTransactionById(db.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found", "reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"receipt": receipt})
}