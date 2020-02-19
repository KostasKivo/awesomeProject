package main

import (
	"awesomeProject/libBooks"
	"fmt"
)

const MENU string = "Welcome to this online library\n" +
	"Press 1 to view available books\n" +
	"Press 2 to add a new book\n" +
	"Press 3 to remove a book"


func main() {

	for{
		fmt.Println(MENU)

		fmt.Println("Please select action!")

		//Gets the user input and stores it in action variable
		var action int
		fmt.Scan(&action)

		//fmt.Printf("This is your action %v" ,action)

		switch action {
		case 1:
			print("List of Available Books" +
				"-----------------------------------\n")
			libBooks.PrintAvailBooks()
		case 2:
			var title string
			var author string

			fmt.Println("Please give the title of the book")
			fmt.Scan(&title)

			fmt.Println("Please give the name of the author of the book")
			fmt.Scan(&author)

			libBooks.AddBook(libBooks.NewBook(title,author))
		case 3:
			fmt.Println("Select the id of the book you want to remove")
			var remId int
			fmt.Scan(&remId)

			libBooks.RemoveBook(remId)
		default:
			print("Please select a supported option")
		} // switch
	} // for loop
}