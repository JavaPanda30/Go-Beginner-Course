package routes

import (
	"database/sql"
	"net/http"

	"example.com/financetracker/db"
	"example.com/financetracker/models"
	"github.com/gin-gonic/gin"
)

func GetAccountByUserID(c *gin.Context) {
	userID := c.Param("user_id")

	query := `SELECT id, user_id, amount FROM accounts WHERE user_id = $1`
	var account models.Account
	err := db.DB.QueryRow(query, userID).Scan(&account.Id, &account.UserID, &account.Amount)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"account": account})
}

func GetAllAccounts(c *gin.Context) {
	query := `SELECT id, user_id, amount FROM accounts`
	rows, err := db.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve accounts", "reason": err.Error()})
		return
	}
	defer rows.Close()

	var accounts []models.Account
	for rows.Next() {
		var account models.Account
		if err := rows.Scan(&account.Id, &account.UserID, &account.Amount); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse account data", "reason": err.Error()})
			return
		}
		accounts = append(accounts, account)
	}

	c.JSON(http.StatusOK, gin.H{"accounts": accounts})
}
