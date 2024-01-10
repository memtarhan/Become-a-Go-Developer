package main

import (
	"fmt"
	"time"
)

func trigger() {
	// use fmt package to print a string
	fmt.Println("Hello world!")

	// use time package to print the current date
	fmt.Printf("Today is %s", time.Now().Weekday())
}
