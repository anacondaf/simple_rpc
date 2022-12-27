package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Book struct {
	Id   string
	Name string
}

func main() {
	client, err := rpc.DialHTTP("tcp", ":3000")

	if err != nil {
		log.Fatal(err)
	}

	var reply = new([]Book)

	if err := client.Call("BookService.GetAllBooks", "Nothing", reply); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", *reply)

	var newBook = Book{Id: "2", Name: "The blizzard"}
	if err := client.Call("BookService.AddBook", newBook, reply); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", *reply)
}
