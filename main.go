package main

import (
	"fmt"
	"strings"

	// import helper package from helper folder we need to specify the path to the package
	"booking_app/helper"
)

var conferenceName string = "2023 Go Conference"

// using := is the same as var conferenceName string = "Conference Name"
const conferenceTickets int = 50

// to make function available to be exported it has to start with capital letter
var RemainingTickets int = 50

// Slice in Go is dynamic array, it can grow and shrink
var ListOfBookings = []string{} // empty slice of strings

// array definition has to have a fixed size [50] and type of data in array e.g. [10]string{} that definition is also its type
// var ListOfBookings = [50]string{} // empty array of 50 strings
// var bookings [50]string // empty array of 50 strings

func main() {

	greetUsers()

	for {

		userFirstName, userLastName, userEmail, numberOfTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(userFirstName, userLastName, userEmail, numberOfTickets, RemainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(numberOfTickets, userFirstName, userLastName)

			firstNames := getFirstNames()
			fmt.Printf("People registered are: %v\n", strings.Join(firstNames, ", "))

			var noTicketsRemaining bool = RemainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("The conference is now fully booked.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("Please enter a valid name.")
			}
			if !isValidEmail {
				fmt.Println("Please enter a valid email.")
			}
			if !isValidTicketNumber {
				fmt.Printf("Please enter a valid number of tickets. There are only %v tickets remaining.\n", RemainingTickets)
			}
		}
	}

}

func greetUsers() {
	fmt.Printf("Conference name is %v \n", conferenceName)
	fmt.Printf("Welcome to conference! There are still %v tickets available.\n", RemainingTickets)
}

func getFirstNames() []string {
	// iteration over the slice, we use _ do discard the index varable which has to be used in the loop
	// in function definition we need to specify the type of the variable we are iterating over and what type of data we are returning
	firstNames := []string{}

	for _, booking := range ListOfBookings {
		// slice string on ' ' and return array of strings
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var userFirstName string
	var userLastName string
	var userEmail string
	var numberOfTickets int

	fmt.Println("What is your first name?")
	// &userName is a pointer to the userName variable, pointers are also called special variables in GO
	fmt.Scanln(&userFirstName)

	fmt.Println("What is your Last name?")
	// &userName is a pointer to the userName variable, pointers are also called special variables in GO
	fmt.Scanln(&userLastName)

	fmt.Println("What is your email?")
	// &userName is a pointer to the userName variable, pointers are also called special variables in GO
	fmt.Scanln(&userEmail)

	fmt.Println("How many tickets do you want to book?")
	fmt.Scanln(&numberOfTickets)

	return userFirstName, userLastName, userEmail, numberOfTickets
}

func bookTicket(numberOfTickets int, userFirstName string, userLastName string) {
	// save name to slice
	RemainingTickets = RemainingTickets - numberOfTickets
	ListOfBookings = append(ListOfBookings, userFirstName+" "+userLastName)

	fmt.Printf("There are now %v tickets remaining.\n", RemainingTickets)
	fmt.Printf("There are %v people registered for the conference.\n", len(ListOfBookings))
}
