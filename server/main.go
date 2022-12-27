package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

var books = []Book{
	{Id: "1", Name: "The dragon"},
}

type Book struct {
	Id   string
	Name string
}

func (b *Book) GetAllBooks(nothing string, reply *[]Book) error {
	*reply = books
	return nil
}

func main() {
	if err := rpc.RegisterName("BookService", new(Book)); err != nil {
		log.Fatal(err)
	}

	rpc.HandleHTTP()

	port := ":3000"
	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("RPC server is serve over http connection at %v", port)

	if err := http.Serve(listener, nil); err != nil {
		log.Fatal(err)
	}
}
