package chatController

import (
	"controllers/userController"
	"net"
)

type Chat struct {
	Users    []*userController.User
	Connect  chan net.Conn
	Outgoing chan string
}
