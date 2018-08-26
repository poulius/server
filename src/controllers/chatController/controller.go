package chatController

import (
	"controllers/userController"
	"fmt"
	"net"
)

func StartChat(listener net.Listener) {
	chat := &Chat{
		Users:    make([]*userController.User, 0),
		Connect:  make(chan net.Conn),
		Outgoing: make(chan string),
	}

	chat.listen()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		chat.Connect <- conn
	}
}

func (chat *Chat) listen() {
	go func() {
		for {
			select {
			case msg := <-chat.Outgoing:
				chat.broadcast(msg)
			case conn := <-chat.Connect:
				chat.join(conn)
			}
		}
	}()
}

func (chat *Chat) join(conn net.Conn) {
	user := userController.NewUser(conn)
	chat.Users = append(chat.Users, user)
	go func() {
		for {
			chat.Outgoing <- <-user.Incoming
		}
	}()
}

func (chat *Chat) broadcast(data string) {
	msg := fmt.Sprintf("%s", data)
	for i, user := range chat.Users {
		if !user.Status {
			chat.Users = append(chat.Users[:i], chat.Users[i+1:]...)
		}
		user.Outgoing <- msg
	}
}
