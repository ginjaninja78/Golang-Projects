package helper

import (
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remTickets uint) (bool, bool, bool) {
	// Demographic input validation
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketQty := userTickets > 0 && userTickets <= remTickets
	return isValidName, isValidEmail, isValidTicketQty
}

// Function to display message in popup
// func displayMessageInPopup(message string) {
// 	err := beeep.Alert("Ticket Confirmation", message, "")
// 	if err != nil {
// 		fmt.Println("Failed to display popup:", err)
// 	}
// }
// Call the function inside sendTicket
// displayMessageInPopup(ticket)

// func displayPopup(message string) {
// 	cmd := exec.Command("osascript", "-e", fmt.Sprintf(`display notification "%s" with title "Ticket Notification"`, message))
// 	err := cmd.Run()
// 	if err != nil {
// 		fmt.Println("Failed to display popup:", err)
// 	}
// }

// Call the displayPopup function in the main function after sending the ticket
// displayPopup(ticket)
