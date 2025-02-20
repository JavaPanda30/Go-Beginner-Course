package routes

import "github.com/gin-gonic/gin"

func Routes(server *gin.Engine) {
	//Income Routes
	server.POST("/income", AddIncome)
	server.GET("/income", GetAllIncome)
	server.GET("/income/:id", GetIncomeById)
	server.PUT("/income/:id", UpdateIncomeRecord)
	server.DELETE("/income/:id", DeleteIncome)

	//Expense Routes
	server.POST("/expense", AddExpense)
	server.GET("/expense", GetAllExpense)
	server.GET("/expense/:id", GetExpenseById)
	server.PUT("/expense/:id", UpdateExpenseRecord)
	server.DELETE("/expense/:id", DeleteExpense)

	//Final Balance Sheet
	server.GET("/final", GetBalance)
}

// Extras

// Income Filters & Reports

// GET /income?category=salary → Filter income by category
// GET /income?month=02&year=2025 → Get monthly income
// GET /income/summary → Get total income summary

// Expense Filters & Reports

// GET /expenses?category=food → Filter by category
// GET /expenses?month=02&year=2025 → Get monthly expenses
// GET /expenses/summary → Get total spending
