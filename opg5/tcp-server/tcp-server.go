package main

import (
	"encoding/json"
	"fmt"
	"net"
)

var info = message{"Navn", "navn@email.com"}

type message struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func handler(c net.Conn) {
	// m := message{"Fredrik", "fredm18@uia.no"}
	b, err := json.Marshal(info)

	if err != nil {
		// Handle error
	}

	c.Write([]byte(b))
	c.Close()
}

func main() {
	fmt.Println("SERVER LAUNCHING")
	l, err := net.Listen("tcp", ":5000")

	if err != nil {
		fmt.Println("Error listening: ", err.Error())
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
		}
		go handler(c)
	}
}
