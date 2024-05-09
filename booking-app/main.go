package main

import (
	"fmt"
	"strings"
	"github.com/ginjaninja78/Golang-Projects/booking-app/helper"
)

// Package level variables, requires full definition syntax
const confTickets int = 50

var remTickets uint = 50
var confName string = "Go Conference"
var bookings = []string{} // Bookings slice

// "main" function is the default function go looks for to run. Additional functions must be explicitly called. Usage: functionName()
func main() {

	greetUsers()

	for len(bookings) < 50 && remTickets > 0 {

		// isInvalidCity := city != "Singapore" && city != "London"
		// Input validation for ticket quantity
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketQty := helper.ValidateUserInput(firstName, lastName, email, userTickets, remTickets)

		if isValidName && isValidEmail && isValidTicketQty {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of the bookings are: %v\n", firstNames)

			if remTickets == 0 {
				// End program
				fmt.Println("Conference is sold out, see you next year.")
				break
			}
		} else {
			// Improve error handling logic
			if !isValidName {
				fmt.Println("The name you entered was too short. Please try again.")
			}
			if !isValidEmail {
				fmt.Println("Email format entered is invalid. Please try again.")
			}
			//Look into variable parsing order issues
			if !isValidTicketQty {
				fmt.Printf("Only %v tickets remain, please try again.\n", remTickets)
			}
		}

	}

}

// *when using package level variables you don't need to pass them into the function separately ie. greetUsers(cName string, cTickets int, rTickets uint)
func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\nWe have %v tickets and %v tickets remain.\nGet your tickets here to attend.\n", confName, confTickets, remTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	// Variable declaration
	var firstName string
	var lastName string
	var userTickets uint
	var email string
	// var city string

	// Ask user for their Name, email and ticket quantity
	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email:")
	fmt.Scan(&email)
	// isValidEmail := strings.Contains(email, "@")
	// fmt.Printf("isValidEmail: %v\n", isValidEmail)

	fmt.Println("Please enter the number of tickets you wish to purchase:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remTickets = remTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you, %v %v for purchasing %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining for %v\n", remTickets, confName)
}
