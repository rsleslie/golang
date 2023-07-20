package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Contact struct{
	Name 			string
	LastName 		string
	NickName 		string
	PhoneNumber		int
	DarkestSecret	string
}

func createContact()Contact{
	var scanner *bufio.Scanner 
	var newContact Contact
	scanner = bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the name: ")
	fmt.Scan(&newContact.Name)
	fmt.Print("Enter the last name: ")
	fmt.Scan(&newContact.LastName)
	fmt.Print("Enter the nickname: ")
	fmt.Scan(&newContact.NickName)
	fmt.Print("Enter the phone number: ")
	fmt.Scan(&newContact.PhoneNumber)
	fmt.Print("Enter the Darkest secret: ")
	fmt.Scan(&newContact.DarkestSecret)
	scanner.Scan()
	newContact.DarkestSecret = strings.TrimSpace(scanner.Text())

	return newContact 
}