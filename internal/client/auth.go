package client

import "fmt"

func AuthUser() (string, bool) {

	// there should be databases here

	var username, password string

	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	if username == "admin" && password == "12345" {
		return username, true
	}

	return "", false

}
