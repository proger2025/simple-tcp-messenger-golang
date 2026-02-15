package main

import (
	"encoding/json"
	"fmt"
	"golang-messenger/internal/models"
	"log"
	"net"
	"sync"
)


var clients = make(map[string]net.Conn)

const (
	address = ":8080"
)

func main () {
	var mx sync.Mutex
	l, err := net.Listen("tcp", address)

	log.Println("Server start. Port: ", address)

	if err != nil {
		log.Println("Failed. Error:", err)
	}


	for {
		conn, err := l.Accept()
		
		if err != nil {
			log.Println("Disconnect")
			continue
		}

		go readLoop(conn, &mx)
	}

	
	

	
	
}


func readLoop(conn net.Conn, mx *sync.Mutex) {
	var username string

	defer func() {
		mx.Lock()
		delete(clients, username)
		mx.Unlock()
		conn.Close()
	}()

	for {
		var msg models.Message	

		json.NewDecoder(conn).Decode(&msg)

		// add conn

		if msg.Type == models.AuthType {
			mx.Lock()

			clients[msg.Sender] = conn
			username = msg.Sender

			mx.Unlock()
			log.Println(msg.Sender,"has joined")

		} else if msg.Type == models.ChatType {
			mx.Lock()
			for _, v := range clients{
				if v != conn {
					json.NewEncoder(v).Encode(msg)
				}
				
				
			}
			mx.Unlock()
		}
		

		fmt.Println(msg)
	}
}





