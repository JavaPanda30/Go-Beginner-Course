package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	note "example.com/notesapp/Note"
	todo "example.com/notesapp/Todo"
)

// a certain value has a certain method so whichever struct has implements interface will always have the Save function that reture an error.

// usual naming = function + 'er'

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputable interface {
	saver
	displayer
}

func outputData(data outputable) error {
	data.Display()
	return SaveData(data)
}

func SaveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Print(err)
		return err
	} else {
		fmt.Print("Saved Successfully")
		return nil
	}
}

func GetTodoData() (string, error) {
	return GetUserInput("Todo Content: ")
}

func GetNoteData() (string, string, error) {
	title, err := GetUserInput("Note title: ")
	if err != nil {
		return "", "", err
	}
	content, err := GetUserInput("Note Content: ")
	if err != nil {
		return "", "", err
	}
	return title, content, err
}

func GetUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	// var value string
	// fmt.Scan(&value)
	// if value == "" {
	// 	return "", errors.New("error :: empty value")
	// }
	// return value, nil
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("invalid input")
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text, nil
}

//

// if we want any type of value

func printSomething(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Integer Value Printed")
	case bool:
		fmt.Println("Boolean Value Printed")
	case string:
		fmt.Println("String Value Printed")
	default:
		fmt.Println("Maybe Something Wierd Printed")
	}
	fmt.Println(value)
}

func main() {
	printSomething(1)
	printSomething("Hello Bye")
	printSomething(true)
	printSomething(12.34)

	title, content, _ := GetNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Print("Error!! Invalid Input")
		return
	}

	text, err := GetTodoData()
	if err != nil {
		fmt.Println("Error")
		return
	}

	userTodo := todo.Todo{Text: text}

	// userTodo.Display()
	// err = userTodo.Save()
	// if err != nil {
	// 	fmt.Print(err)
	// 	return
	// }
	// fmt.Println("Saved Successfully")

	//  this can be replaced by interface ->

	if SaveData(userTodo) != nil {
		return
	}

	// and now this can be replaced by embedded interfaces
	outputData(userTodo)
	outputData(userNote)
}
