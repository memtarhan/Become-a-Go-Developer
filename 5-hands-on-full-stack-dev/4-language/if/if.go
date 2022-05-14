package main

import "fmt"

// The if statement

func getNumber() int { return 10 }

func main() {
	number := getNumber()

	if number == 10 {
		fmt.Println("Number is 10")
	}

	// or
	if x := getNumber(); x == 10 {
		fmt.Println("x is 10")
	}

	if number == 5 {
		fmt.Println("Number is 5")
	} else {
		fmt.Println("Number is not 5")
	}

	if number == 5 {
		fmt.Println("Number is 5")

	} else if number > 10 {

	} else {
		
	}
}
