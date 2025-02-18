package main

import (
	"fmt"
	"strings"
)

func main() {

	//maps

	mp := make(map[string]int)

	mp["k1"] = 7
	mp["k2"] = 13
	fmt.Println("map: ", mp)

	v2 := mp["k3"]
	fmt.Println(v2)

	fmt.Println(len(mp))

	delete(mp, "k2")
	fmt.Println(mp)

	clear(mp)
	fmt.Println(mp)

	_, prs := mp["k1"]
	fmt.Println("prs", prs)
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)

	//Ranges

	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	for i, num := range nums {
		if num == 3 {
			fmt.Println(i)
		}
	}

	fmt.Println(sum)

	var name = "Suyash"

	//prints each character
	fmt.Printf("%c", name[0])
	fmt.Print(strings.Fields(name))

}
