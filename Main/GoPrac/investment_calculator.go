package main

import (
	"fmt"
	"math"
	// "strconv"
	// "log"
)

func calculate_Investment() {
	var investmentAmount, years, expectedReturnRates, inflationRate float64
	fmt.Scan(&investmentAmount)
	fmt.Scan(&years)
	fmt.Scan(&expectedReturnRates)
	fmt.Scan(&inflationRate)
	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRates/100, float64(years))

	fututeInflationValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(fututeInflationValue)
	// writeInFile(strconv.FormatFloat(futureValue, 'f', 2, 64))
	// message, err := getDataFromFile()
	// fmt.Print(message)
	// if err!=nil {
	// 	log.Fatal(err)
	// }
}

// Key difference in const and  that var is const can be reassigned.
// now if i want to take user input
