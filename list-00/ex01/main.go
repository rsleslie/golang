package main

import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	reander := bufio.NewReader(os.Stdin)
	var option string
	var oldIndex int = 0
	var contacts []Contact
	// var counterContcts int = 0

	for {
		fmt.Print("You can choose ADD contact, SEARCH contact or EXIT to exit the program: ")
		option, _ = reander.ReadString('\n')
		
		if option == "ADD\n" {
			if len(contacts) <= 8 {
				contacts = append(contacts, createContact())
			}else{
				contacts[oldIndex] = createContact()
				oldIndex = (oldIndex + 1) % 7
			}
		}else if option == "SEARCH\n" {
			search(contacts)
		}else if option == "EXIT\n" {
			return
		}else {
			fmt.Println("Invalid option")
		}
	}
}