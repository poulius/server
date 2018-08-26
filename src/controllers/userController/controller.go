package userController

import (
    "bufio"
    "net"
    "fmt"
    "encoding/json"
)

func NewUser(conn net.Conn) *User {
	writer := bufio.NewWriter(conn)

	user := &User{
		Conn:       conn,
		Writer:     writer,
		Outgoing:   make(chan string),
		Incoming:   make(chan string),
		Disconnect: make(chan bool),
		Status:     true,
	}

	go user.write()
	go user.read()

	return user
}

func (user *User) write() {
	for {
		select {
		case <-user.Disconnect:
			user.Status = false
			break
		default:
			msg := <-user.Outgoing
			user.Writer.WriteString(msg)
			user.Writer.Flush()
		}
	}
}

func (user *User) read() {
  var msg Message
	for {
    decoded := json.NewDecoder(user.Conn)

    err := decoded.Decode(&msg)
    if err != nil {
      	user.Incoming <- fmt.Sprintf("%s is offline\n", user.Name)
  			user.Status = false
  			user.Disconnect <- true
  			user.Conn.Close()
  			break
    }

    switch {
		case msg.MessageType == "name":
			user.Name = msg.MessageText
			user.Incoming <- fmt.Sprintf("server: %s is online\n", msg.MessageText)
		case msg.MessageType == "message":
			user.Incoming <- fmt.Sprintf("%s: %s \n", user.Name, msg.MessageText)
    default:
      break
		}
	}
}
