package main

import (
	"errors"
	"fmt"
)

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		panic("Invalid Revenue")
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		panic("Invalid Expense")
	}
	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		panic("Invalid Tax Rate")
	}
	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\n Ratio: %.3f ", ebt, profit, ratio)
	writeInFile(results)
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("invalid input")
	}
	return userInput, nil
}
