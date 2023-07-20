package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func search(contacts[] Contact){
	var name string

	fmt.Print("Whats is the name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	for i := 0; i < len(contacts); i++ {

		if contacts[i].Name == name{
			fmt.Println("Name :", contacts[i].Name)
			fmt.Println("Last name :", contacts[i].LastName)
			fmt.Println("Nickname :", contacts[i].NickName)
			fmt.Println("Phone number :", contacts[i].PhoneNumber)
			fmt.Println("Darkest secret :", contacts[i].DarkestSecret)
			return
		}
	}
	fmt.Println("Contact not found.")
}

func add(newContact Contact, contacts []Contact) {
	if len(contacts) < 8 {
		contacts = append(contacts, newContact)
	} else {
		contacts = append(contacts[1:], newContact) 
	}
}