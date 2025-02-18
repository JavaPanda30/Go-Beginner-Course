package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string,error) {
	if name==""{
		return "",errors.New("empty name")
	}
	message := fmt.Sprintf("Hi %v. Welcome Here", name)
	return message,nil
}


func Abc(){
	fmt.Println("Greeted ADC")
}