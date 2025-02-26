package routes

import "github.com/gin-gonic/gin"

func Routes(server *gin.Engine) {
	//Income Routes
	server.POST("/income", AddIncome)
	server.GET("/income", GetAllIncome)
	server.GET("/income/:income_id", GetIncomeById)
	server.PUT("/income/:income_id", UpdateIncomeRecord)
	server.DELETE("/income/:income_id", DeleteIncome)

	//Expense Routes
	server.POST("/expense", AddExpense)
	server.GET("/expense", GetAllExpense)
	server.GET("/expense/:expense_id", GetExpenseById)
	server.PUT("/expense/:expense_id", UpdateExpenseRecord)
	server.DELETE("/expense/:expense_id", DeleteExpense)

	//User Routes
	server.POST("/user", AddUser)
	server.GET("/users", GetAllUser)
	server.GET("/user/:user_id", GetUserById)
	server.PUT("/user/:user_id", UpdateUserRecord)
	server.DELETE("/user/:user_id", DeleteUser)

	//Filters and Final Balance Sheet for income expense
	server.GET("/user/balance/:user_id", GetBalance)
	server.GET("/user/income/:user_id", GetTotalIncomeOfUser)
	server.GET("/user/expense/:user_id", GetTotalExpenseOfUser)

	//User Transfer
	server.POST("/transfer", SendMoney)
	server.GET("/account/:user_id", GetAccountByUserID)
	server.GET("/accounts", GetAllAccounts)
	server.GET("/receipt/:transaction_id", GetReceipt)
}
