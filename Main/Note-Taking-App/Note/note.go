package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedOn time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("empty fields")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedOn: time.Now(),
	},nil
}


func (note Note) Display(){
	fmt.Print("Note Title: ")
	fmt.Println(note.Title)
	fmt.Print("Your Note has Following Content: ")
	fmt.Println(note.Content)
}

func (note Note) Save() error{
	fileName:=strings.ReplaceAll(note.Title," ","_")
	fileName =strings.ToLower(fileName) + ".json"
	json,err:=json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName,json,0644)
}

