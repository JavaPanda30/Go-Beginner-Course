package main

import "fmt"

func main() {
	result := add(12.5, 21)
	fmt.Println(result)
}

// this T means i can use any of the following specified type in input or instead can give any or interface{} that means a interface with no function in which all the criteria are satisfied and i can call any type of data in this category

func add[T int | float64 | string](a, b T) T {
	return a + b
}
