package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golang-messenger/internal/client"
	"golang-messenger/internal/models"
	"log"
	"net"
	"os"

)

var users = make(map[string]net.Conn)

func main() {
	conn, err := net.Dial("tcp", ":8080")

	if err != nil {
		log.Println("Connetcion failed")
	}

	username, auth := client.AuthUser()
	if auth {
		msg := models.Message {
			Type: "AUTH",
			Sender: username,
			Content: "new client",
		}

		b, _ := json.Marshal(msg)


		conn.Write(b)
	}

	
	go readLoop(conn)
	go writeLoop(conn, username)

	select {}

}


func readLoop(conn net.Conn) {
	for {
		var msg models.Message

		err := json.NewDecoder(conn).Decode(&msg)

		if err != nil {
			log.Println(err)
		}
		if msg.Type == models.ChatType {
			fmt.Printf("[%v]: %v \n\n", msg.Sender, msg.Content)
		}
		
	


	}
}


func writeLoop(conn net.Conn, username string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024), 1024*1024)

	// skip the first input to avoid duplicate ENTER
	scanner.Scan()

	for {
		// fmt.Print("\nENTER: ")
		if !scanner.Scan(){
			return
		}
		
		enter := scanner.Text()

		if enter == "" { continue }

		newMessage := models.Message{
			Type: "CHAT",
			Sender: username,
			Content: enter,
		}

		err := json.NewEncoder(conn).Encode(newMessage)
		if err != nil {
			log.Println("Failed. Error:", err)
		}

	}
}




