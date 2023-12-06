package model

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Clients map[net.Conn]Client
	Mu      sync.Mutex
}

func (s *Server) Handle(conn net.Conn) {
	decoder := json.NewDecoder(conn)
	defer func() {
		operation := func() {
			delete(s.Clients, conn)
		}
		s.SafeOperation(operation)
		err := conn.Close()
		if err != nil {
			fmt.Printf("Error while closing the connection: %v\n", err)
		}
	}()

	for {
		var client Client
		var statusType ClientStatusType
		if err := decoder.Decode(&client); err != nil {
			fmt.Printf("JSON Decode error: %v\n", err)
			return
		}

		switch {
		case ClientExitType{}.checkStatus(client.Message):
			statusType = ClientExitType{}
		case ClientEnterType{}.checkStatus(client.Message):
			statusType = ClientEnterType{}
		case ClientSendDataType{}.checkStatus(client.Message):
			statusType = ClientSendDataType{}
		default:
			fmt.Println("Undefined status type of client message")
			continue
		}

		fmt.Println(statusType.getMessage(client))
	}
}

func (s *Server) SafeOperation(operation func()) {
	s.Mu.Lock()
	operation()
	s.Mu.Unlock()
}
