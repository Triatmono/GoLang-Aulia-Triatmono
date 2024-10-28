package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}
 
	fmt.Printf("Welcome to %v booking applications\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket now")

	for remainingTickets > 0 && len(bookings) < 50 {
		
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		
		// ask user for their name
		fmt.Println("Please enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email:")
		fmt.Scan(&email)

		fmt.Println("Please enter number of tickets:")
		fmt.Scan(&userTickets)
		
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber{
			
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings,firstName + " " + lastName )

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings{
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining{
				fmt.Println("Tickets are sold out. Come back another time.")
				break
			}
		}  else {
			if !isValidName{
				fmt.Println("First name and last name must be at least 2 characters.")
			} 
			if !isValidEmail{
				fmt.Println("Email must be valid.")
			}
			if  !isValidTicketNumber{
				fmt.Println("Number of tickets is invalid.")
			}			
		}
	}

	city := "London"

	switch city {
		case "New York":
			//code
		case "Singapore":

		case "London":

		case "Berlin":

		case "Hong Kong":
		
		default:
			fmt.Print("No valid city selected\n")
	}
}
