package main

import (
	"GoSocketChatApp/model"
	"fmt"
	"net"
)

const (
	networkType        = "tcp"
	networkHost        = "localhost"
	defaultNetworkPort = "3000"
)

func main() {
	ln, err := net.Listen(networkType, fmt.Sprintf("%v:%v", networkHost, defaultNetworkPort))
	if err != nil {
		fmt.Println(err)
		return
	}

	server := &model.Server{
		Clients: make(map[net.Conn]model.Client),
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		server.SafeOperation(func() {
			server.Clients[conn] = model.Client{Name: "Anonymous"}
		})

		go server.Handle(conn)
	}
}
