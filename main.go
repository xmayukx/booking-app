package main

import (
	"booking-app/helper"
	"fmt"

	// "sync"
	"time"
)

var confName = "Go Conf"

const confTix uint = 50

var remaininTix uint = confTix
var bookings = make([]UserData, 0)

type UserData struct {
	fname   string
	lname   string
	email   string
	userTix uint
}

// var wg = sync.WaitGroup{}

func main() {

	greet()

	for {
		fname, lname, email, userTix := getuserInput()
		//validate
		isValidEmail, isValidUserTix, isValidName := helper.ValidateUserInput(fname, lname, email, userTix, remaininTix)

		if isValidEmail && isValidName && isValidUserTix {

			bookings, remaininTix = bookTix(userTix, fname, lname, email)
			// wg.Add(1)
			go sendTix(userTix, fname, lname, email)
			fmt.Printf("All the bookings are: %v ", getFirstNames())

			noTixAvailable := remaininTix == 0
			if noTixAvailable {
				//end prog
				fmt.Println("Our conference is booked out. Come back next year!")

			}
		} else {

			if !isValidName {
				fmt.Print("The name you have entered is invalid\n")
			}
			if !isValidEmail {
				fmt.Print("The email you have entered doesn't contains '@' symbol.\n")
			}
			if !isValidUserTix {
				fmt.Printf("We only have %v available. You can't book %v tickets.\n", remaininTix, userTix)
			}
		}
	}
	// wg.Wait()
}

func greet() {
	fmt.Printf("confTix is %T, remaininTix is %T and confName is %T.\n", confTix, remaininTix, confName)
	fmt.Printf("Welcome to %s booking application.\n", confName)
	fmt.Printf("We have total of %v tickets and %v tickets are still available.\n", confTix, remaininTix)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.fname)
	}
	return firstNames
}

func getuserInput() (string, string, string, uint) {
	var fname string
	var lname string
	var email string
	var userTix uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&fname)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lname)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTix)

	return fname, lname, email, userTix
}

func bookTix(userTix uint, fname string, lname string, email string) ([]UserData, uint) {
	remaininTix = remaininTix - userTix

	var userData = UserData{
		fname:   fname,
		lname:   lname,
		email:   email,
		userTix: userTix,
	}

	bookings = append(bookings, userData)

	fmt.Printf("list of bookings is %v.\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", fname, lname, userTix, email)
	fmt.Printf(" %v are left for %v. \n", remaininTix, confName)
	return bookings, remaininTix
}

func sendTix(userTix uint, fname string, lname string, email string) {
	time.Sleep(10 * time.Second)
	var tix = fmt.Sprintf("%v tickets for %v %v", userTix, fname, lname)
	fmt.Println("########################################")
	fmt.Printf(" Sending ticket %v to email %v\n", tix, email)
	fmt.Println("########################################")
	// wg.Done()
}
