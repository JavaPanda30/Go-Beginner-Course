package main

import "fmt"

func UseMake() {
	//make is used to create a slice of any type that will be having a variable length
	name := make([]string, 2)
	name[0] = "Suyash"
	name[1] = "Yash"
	name = append(name, "George")
	name = append(name, "Ram") //not in the capacity but added at the end of the two length no matter they are empty or not
	fmt.Println(name)

	//Problem with making an empty initialization then for every addition go will have to reallocate memory and thus make is used so that it will have a basic rough estimate of size and if that gets full only after that GO assigns new value so it is not like array where size is fixed and better than slice where each addition costs computation
}
