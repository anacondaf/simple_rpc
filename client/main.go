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

	var reply []Book

	if err := client.Call("BookService.GetAllBooks", "Nothing", &reply); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", reply)
}
