package main

import (
	"fmt"
	"time"
)

var ConferenceName = "Go Conference"

const ConferenceTickets = 50

var RemainingTickets = 50
var bookings = make([]UserData, 0)

type UserData struct {
	fname          string
	lname          string
	email          string
	numberOfTicket int
}

func main() {

	greetuser(ConferenceName, ConferenceTickets, RemainingTickets) //call for the function greetuser

	//array
	//var bookings [50]string
	//slicing
	// for printing statement
	//infinte for loop
	for {
		//to check the datatype
		//fmt.Printf("ConferenceName is %T ConferenceTickets is %T RemainingTickets %T \n", ConferenceName, ConferenceTickets, RemainingTickets)

		fname, lname, usertickets, email := getUserInput()
		isValindName, isValindEmail, isValindTicketNumber := ValidateUserInput(fname, lname, email, usertickets, RemainingTickets)

		if isValindName && isValindEmail && isValindTicketNumber {
			RemainingTickets, bookings = BookTickets(usertickets, RemainingTickets, bookings, fname, lname, email)
			sendTickets(usertickets, fname, lname, email)
			printFirstName(bookings)

			//if condition
			if RemainingTickets == 0 {
				fmt.Print("Oppp's Sorry we are sold out!!!! \n")
				break
			}
		} else {
			if !isValindName {
				fmt.Print("Please enter valid First and Last name \n")
			}
			if !isValindEmail {
				fmt.Print("Please enetr valid Email id \n")
			}
			if !isValindTicketNumber {
				fmt.Print("Please enter valid Ticket value \n")
			}

		}
	}

}
func greetuser(ConferenceName string, RemainingTickets int, ConferenceTickets int) {
	fmt.Printf("Welcome to  %v booking application \n", ConferenceName)
	fmt.Println(" book your tickets here")
	fmt.Printf("we have total of  %v tickets and we have %v remaining tickets\n", ConferenceTickets, RemainingTickets)

}

func printFirstName(bookings []UserData) {
	var firstnames []string
	for _, booking := range bookings {
		firstnames = append(firstnames, booking.fname)
	}
	fmt.Printf("The first names of bookings are: %v\n", firstnames)
}

func getUserInput() (string, string, int, string) {
	var fname string
	var lname string
	var usertickets int
	var email string

	//taking input from user
	fmt.Printf("enter your first name \n")
	fmt.Scan(&fname)

	fmt.Printf("enter your last name \n")
	fmt.Scan(&lname)

	fmt.Printf("enter your number of tickets \n")
	fmt.Scan(&usertickets)

	fmt.Printf("enter your Email id \n")
	fmt.Scan(&email)
	//it is for array
	//bookings[0] = username
	return fname, lname, usertickets, email
}

func BookTickets(usertickets int, RemainingTickets int, bookings []UserData, fname string, lname string, email string) (int, []UserData) {

	//for slicing
	RemainingTickets = RemainingTickets - usertickets

	// create a map for a use

	var userData = UserData{
		fname:          fname,
		lname:          lname,
		email:          email,
		numberOfTicket: usertickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of booking is %v\n", bookings)
	//fmt.Printf("the whole array %v \n", bookings)
	//fmt.Printf("the first element array %v \n", bookings[0])
	//fmt.Printf("the  array type %T \n", bookings)
	//fmt.Printf("the length of array %v\n", len(bookings))
	fmt.Printf("user %v booked %v tickets \n", fname+" "+lname, usertickets)
	fmt.Printf("Thanksss %v for booking %v tickets \n", fname, usertickets)
	fmt.Printf("the customers email id is %v \n", email)
	fmt.Printf("the remaining tickets are %v HURRY UP! \n", RemainingTickets)
	return RemainingTickets, bookings

}

func sendTickets(usertickets int, fname string, lname string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", usertickets, fname, lname)
	fmt.Println("###############")
	fmt.Printf("sending tickets: \n %v \n to email address%v \n", ticket, email)
	fmt.Println("###############")
}
