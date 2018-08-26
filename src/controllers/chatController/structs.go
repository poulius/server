package chatController

import (
  "net"
  "controllers/userController"
)

type Chat struct {
	Users  []*userController.User
	Connect  chan net.Conn
	Outgoing chan string
}
