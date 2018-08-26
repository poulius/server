package main

import (
	"controllers/chatController"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Launching server...")

	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Println(err)
	}

	chatController.StartChat(listener)
}
