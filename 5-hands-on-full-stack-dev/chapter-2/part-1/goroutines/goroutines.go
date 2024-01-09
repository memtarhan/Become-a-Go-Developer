package main

import (
	"fmt"
	"time"
)

// Goroutines

/*
	A goroutine is simply defined as a light-weight thread that you can use in your
	program; it's not a real thread. In Go, when you define a piece of code as a new
	goroutine, you basically tell the Go runtime that you would like this piece of
	code to run concurrently with other goroutines.

	Every function in Go lives in some goroutine. For example, the main function
	that we discussed in the previous chapter, which is usually the entry point
	function for your program, runs on what is known as the main goroutine.

	So, how do you create a new goroutine? You just append the go keyword before the
	function that you would like to run concurrently

	go someFunction() { }
*/

func runSomeLoop(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Printing:", i)
	}
}

func main() {
	go runSomeLoop(10)
	// block the main goroutine for 2 seconds
	time.Sleep(2 * time.Second)
	fmt.Println("Hello, playground")
}
