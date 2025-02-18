package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (text Todo) Display() {
	fmt.Println("Todo Text: ")
	fmt.Println(text.Text)
}

func (text Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(text)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("empty field")
	}
	return Todo{
		Text: text,
	}, nil
}
