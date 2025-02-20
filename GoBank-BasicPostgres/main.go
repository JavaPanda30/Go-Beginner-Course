package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
		return
	}
	err=store.DbInit()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%+v\n", store)
	server := NewAPIServer(":3000")
	server.run()
}
