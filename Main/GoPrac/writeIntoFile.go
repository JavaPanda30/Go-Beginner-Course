package main

import (
	// "errors"
	"fmt"
	"os"
	// "strconv"
)

func writeInFile(message string) {
	os.WriteFile("goWrite.txt", []byte(message), 0644)
	fmt.Println("File Created and Written")
}

// func getDataFromFile() (float64, error) {
// 	fmt.Println("Get File func called")
// 	data, err := os.ReadFile("goWrite.txt")
// 	if err == nil {
// 		datawritten := string(data)
// 		num, _ := strconv.ParseFloat(datawritten, 64)
// 		fmt.Println(num)
// 		return num, nil
// 	} else {
// 		panic("Faied to Fetch!!")
// 		return 0, errors.New("failed to fetch the file")
// 	}
// }
