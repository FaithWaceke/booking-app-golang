package main

import "fmt"

func main() {
	var confrenceName = "Go confrence"
	const confrenceTicket = 100
	var remainingTicket = 100

	fmt.Printf("Welcome to %v booking application\n", confrenceName)
	fmt.Printf("We have a total of %v ticket and %v are still available.\n", confrenceTicket, remainingTicket)
	fmt.Printf("Get your ticket here to attend\n")

	var firstName string
	var lastName string
	var email string
	var userTicket int
	// ask user for their name
	fmt.Println("Enter your firstName")
	fmt.Scan(&firstName)

	fmt.Println("Enter your lastName")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets")
	fmt.Scan(&userTicket)

	remainingTicket = remainingTicket - userTicket

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTicket, confrenceName)

}
