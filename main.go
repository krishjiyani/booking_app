package main

import (
	"fmt"
	"strings"
)

func main() {

	const conferenceTickets int = 50
	var remainingTickets uint = 50
	conferenceName := "Go Conference"
	bookings := []string{}
	
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	
	
	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint
		var cityname string

		// asking for user input
		fmt.Println("Enter Your First Name: ")
		fmt.Scanln(&firstName)

		fmt.Println("Enter Your Last Name: ")
		fmt.Scanln(&lastName)

		fmt.Println("Enter Your Email: ")
		fmt.Scanln(&email)
		fmt.Println("Enter Your First city name: ")
		fmt.Scanln(&cityname)

		fmt.Println("Enter number of tickets: ")
		fmt.Scanln(&userTickets)

		var isValidName = len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isvalidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}

		if isValidName && isValidEmail && isvalidTicketNumber {
			// book ticket in system
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("These are all our bookings: %v\n", bookings)
			fmt.Printf("The first names: %v\n", firstNames)
		}



		if remainingTickets == 0 {
			// end the program
			fmt.Println("Our conference is booked out. Come back next year.")
			break // Exit the loop when no tickets remain
		} else if userTickets == remainingTickets {
			fmt.Print("Congratulations, you have booked the last of the tickets!\n")
		}

		if !isvalidTicketNumber {
			fmt.Println("The number of tickets you entered is invalid.")
		}
	} // <- closes the `for` loop
}

func greetUsers(conFname string,confTickets int,remainningTickets uint) {
	var confName string
	fmt.Println("Welcome to %v booking application", confName)
	fmt.Printf("Welcome to %v booking application.\nWe have a total of %v tickets, and %v are still available.\nGet your tickets here to attend\n", confName, confTickets, remainningTickets)
	
}
