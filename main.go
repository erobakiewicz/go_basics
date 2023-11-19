package main

import (
	"fmt"
	"strings"
	"sync"

	// import helper package from helper folder we need to specify the path to the package
	"booking_app/helper"
	"time"
)

var conferenceName string = "2023 Go Conference"

// using := is the same as var conferenceName string = "Conference Name"
const conferenceTickets int = 50

// to make function available to be exported it has to start with capital letter
var RemainingTickets int = 50

// Slice in Go is dynamic array, it can grow and shrink, to init slice of maps we need to use 'make' function
var ListOfBookings = make([]UserData, 0) // empty slice of strings

// array definition has to have a fixed size [50] and type of data in array e.g. [10]string{} that definition is also its type
// var ListOfBookings = [50]string{} // empty array of 50 strings
// var bookings [50]string // empty array of 50 strings

// struct is a data structure that can hold different types of data, it is similar to clas in Python or Java
// to init struct we need to use 'type' keyword
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

// If main treat is finished it will not wait for other threads to finish, to wait for other threads to finish we need to use sync package
var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	userFirstName, userLastName, userEmail, numberOfTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(userFirstName, userLastName, userEmail, numberOfTickets, RemainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		// map is a key value pair data structure, to init map we need to use 'make' function
		var userData = UserData{
			firstName:       userFirstName,
			lastName:        userLastName,
			email:           userEmail,
			numberOfTickets: numberOfTickets,
		}
		//add to map
		// userData["firstName"] = userFirstName
		// userData["lastName"] = userLastName
		// userData["email"] = userEmail
		// userData["numberOfTickets"] = strconv.FormatInt(int64(numberOfTickets), 10)

		bookTicket(numberOfTickets, userFirstName, userLastName, userData)
		// use keyword 'go' to run function in a separate thread and add it to wait group
		wg.Add(1)
		go sendTicekt(userFirstName, userLastName, numberOfTickets, conferenceName)

		firstNames := getFirstNames()
		fmt.Printf("People registered are: %v\n", strings.Join(firstNames, ", "))

		var noTicketsRemaining bool = RemainingTickets == 0
		if noTicketsRemaining {
			fmt.Println("The conference is now fully booked.")
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
	// wait for all threads to finish
	wg.Wait()
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
		firstNames = append(firstNames, booking.firstName)
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

func bookTicket(numberOfTickets int, userFirstName string, userLastName string, userData UserData) {
	// save name to slice
	RemainingTickets = RemainingTickets - numberOfTickets
	ListOfBookings = append(ListOfBookings, userData)

	fmt.Printf("There are now %v tickets remaining.\n", RemainingTickets)
	fmt.Printf("There are %v people registered for the conference.\n", len(ListOfBookings))
}

func sendTicekt(userFirstName string, userLastName string, numberOfTickets int, conferenceName string) {
	time.Sleep(3 * time.Second)
	var ticket = fmt.Sprintf("Hi %v %v, thank you for booking %v tickets for %v", userFirstName, userLastName, numberOfTickets, conferenceName)
	fmt.Printf("Sending ticket: \n %v to %v %v\n", ticket, userFirstName, userLastName)
	// remove thread from wait group
	wg.Done()
}
