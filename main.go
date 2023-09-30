package main

import "fmt"

func main() {
	var confrenceName = "Go confrence"
	const confrenceTicket = 50
	var remainingTicket = 50

	fmt.Printf("Welcome to %v booking application\n", confrenceName)
	fmt.Printf("We have a total of %v ticket and %v are still available.\n", confrenceTicket, remainingTicket)
	fmt.Printf("Get your ticket here to attend\n")

	var userName string
	var userTicket int
	// ask user for their name

	userName = "Phebby"
	userTicket = 2
	fmt.Printf("user %v booked %v tickets.\n", userName, userTicket)

	// Naona una endelea vizuri

}
