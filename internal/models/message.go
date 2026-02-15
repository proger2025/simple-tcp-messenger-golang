package models

const (
	AuthType = "AUTH" // AUTH type denotes a service message indicating that the user is authorized
	ChatType = "CHAT" // CHAT type denotes a message written in a chat
)


type Message struct {
	Type string 
	Sender string 
	Content string 
}


