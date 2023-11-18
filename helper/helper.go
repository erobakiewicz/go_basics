package helper

import (
	"strings"
)

// !!! IMPORTANT !!! to export function from package it has to start with capital letter
func ValidateUserInput(userFirstName string, userLastName string, userEmail string, numberOfTickets int, RemainingTickets int) (bool, bool, bool) {
	// in Go function can return multiple values
	var isValidName bool = len(userFirstName) > 2 && len(userLastName) > 2
	var isValidEmail bool = len(userEmail) > 0 && strings.Contains(userEmail, "@")
	var isValidTicketNumber bool = numberOfTickets > 0 && numberOfTickets <= RemainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
