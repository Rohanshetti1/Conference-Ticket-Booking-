package main

import "strings"

func ValidateUserInput(fname string, lname string, email string, usertickets int, RemainingTickets int) (bool, bool, bool) {
	//for validation
	var isValindName bool = len(fname) >= 3 && len(lname) >= 3
	var isValindEmail bool = strings.Contains(email, "@")
	var isValindTicketNumber bool = usertickets > 0 && usertickets <= RemainingTickets

	return isValindName, isValindEmail, isValindTicketNumber //retur syntax
}
