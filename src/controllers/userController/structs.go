package userController

import (
	"bufio"
	"net"
)

type User struct {
	Name       string
	Conn       net.Conn
	Writer     *bufio.Writer
	Incoming   chan string
	Outgoing   chan string
	Disconnect chan bool
	Status     bool
}

type Message struct {
	MessageType string
	MessageText string
}
