package controllers

import "fmt"

func User_options() int{
	var  user_choice int 
	fmt.Println("\n|-------------------------------------------------------------------------------------|")
	fmt.Println("Hello! Welcome to the library Management system.")
	fmt.Println("What would you like to do today?")
	fmt.Printf("1.Add a new book.\n2.Remove an existing book.\n3.Borrow a book.\n4.Return a book.\n5.List all available books.\n6.List borrowed books by a member.\n7.Exit Program.\n")
	fmt.Println("Interact with the console by typing the number of the option you would like to choose.")
	fmt.Println("|--------------------------------------------------------------------------------------|")
	fmt.Printf("Your choice: ")
	fmt.Scanln(&user_choice)
	return user_choice
}