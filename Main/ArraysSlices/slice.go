package main

import "fmt"

func makeFunc() {

	type Product struct {
		title string
		id    int
		price float32
	}

	hobbies := [3]string{"cricket", "Coding", "Singing"}
	fmt.Println(hobbies)
	// 1.
	fmt.Println(hobbies[0])
	// 2.
	fmt.Println(hobbies[1:3])
	fmt.Println(hobbies[1:])
	// 3.
	hobbyslice := hobbies[0:2]
	hobbyslice1 := hobbies[:2]
	fmt.Println(hobbyslice)
	fmt.Println(hobbyslice1)

	// 4.
	capacity := cap(hobbyslice)
	fmt.Println(capacity)
	hobbyslice = hobbyslice[1:3]
	fmt.Println(hobbyslice)

	//5.
	goals := []string{"LearnGO", "LearnAPI"}
	fmt.Println(goals)

	//6.
	goals[1] = "LearnHTTP"
	fmt.Println(goals)
	goals = append(goals, "LearnAPI's")
	fmt.Println(goals)

	//7.
	products := []Product{
		{"GO GuideBook", 1, 199.99},
		{"GO Advance GuideBook", 2, 299.99},
	}
	fmt.Println(products)
	newProd := Product{"GO Backend Book", 1, 399.99}
	products = append(products, newProd)
	fmt.Println(products)

}

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
