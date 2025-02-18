package main

import (
	"fmt"
	"time"
)

func helper(str string,donechan chan bool) {
	for i := 0; i < 5; i++ {
		fmt.Println(str)
	}
	donechan<-true
}
func slowhelper(str string, donechan chan bool) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(str)
	}
	donechan <- true
	
}

func main() {

	//make channel
	done := make(chan bool)

	go slowhelper("hello",done)
	go helper("00011",done)

	for range done{

	}
	
}
