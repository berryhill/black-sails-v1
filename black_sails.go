package main

import (
	"fmt"

	"github.com/black-sails-v0/session"
)

func main() {
	fmt.Println("Welcome to Black Sails")

	Session := session.NewSession()
	fmt.Println(Session)


}

