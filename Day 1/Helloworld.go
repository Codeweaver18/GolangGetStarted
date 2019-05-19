package main

import (
	"fmt"
	"strings"
	"time"
)

//My first helloword struct
func printHelloWolrd() {
	//Learning Printing Hello world
	fmt.Println("Hello world, welcome to go programing")
	fmt.Println(strings.ToUpper("Hellowolrd from second line"))
}

///Getting user input and printing it on screen
func GettingUerInputsFromConsole(fname string, lname string) (fullname string) {

	fullname = "My Name is " + strings.ToLower(fname) + " " + strings.ToUpper(lname)
	return
}

//strings compare
func StringCompare(fname string, lname string) (result bool) {

	if fname == lname {
		return true
	}
	return
}

//simple interest calculation
func simpleinterest(principal int, rate int, duration int) (result int) {

	var amount = principal * (1 + rate*duration)
	result = amount
	return
}

//Playing with time
func getDateTime() {
	var t = time.Now()
	fmt.Println(t.Day())
}

func main() {
	getDateTime()
	fmt.Println(simpleinterest(100, 20, 5))
	fmt.Println(StringCompare("Monday", "Matsumu"))
	fmt.Println(GettingUerInputsFromConsole("Matsumu, Monday", "garba"))
}
