package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	c, err := net.Dial("tcp", ":5000")

	if err != nil {
		panic(err)
	}

	str, err := bufio.NewReader(c).ReadString('\n')
	fmt.Println(str)
}
