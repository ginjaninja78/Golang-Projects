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
