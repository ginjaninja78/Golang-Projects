package main

import (
	"booking-app/helper"
	"fmt"
	"time"
	// "os/exec" package for linux and mac. Used to display popup
	// "beep" package for windows. Windows only. "github.com/gen2brain/beeep"
)

// Package level variables, requires full definition syntax
const confTickets int = 50

var remTickets uint = 50
var confName string = "Go Conference"

// var bookings = make([]map[string]string, 0) // Bookings slice of maps
var bookings = make([]UserData, 0) // Bookings user data struct

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
	// optIn				bool
}

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
			sendTicket(userTickets, firstName, lastName, email)

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
		firstNames = append(firstNames, booking.firstName)
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

	// Create a struct for user details
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
		// optIn: true,
	}
	// For use with maps
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you, %v %v for purchasing %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining for %v\n", remTickets, confName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("######################")
	fmt.Printf("Sending ticket:\n%v\nto email address %v\n", ticket, email)
	fmt.Println("######################")
}
