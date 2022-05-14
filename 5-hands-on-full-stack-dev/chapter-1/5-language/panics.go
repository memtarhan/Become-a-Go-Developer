package main

import "fmt"

// Panic

/*
	In Go, there is a special built-in function called panic. When you invoke panic
	in your code, your program is interrupted, and a panic message is returned. If a
	panic gets triggered and you don't capture it in time, your program will stop
	execution and will exit, so be very careful when you use a panic.
*/

func panicTest(p bool) {
	if p {
		panic("panic requested")
	}
}

func printEnding(message string) {
	fmt.Println(message)
}

func doSomething() {
	// In here we use the keyword "defer"
	// This will call printEnding() right after doSomething()

	defer printEnding("doSomething() just ended")
	defer printEnding("doSomething() just ended 2")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	/*
		In the preceding code, when we made use of defer, we effectively asked for the
		printEnding() function to be executed right after doSomething() finishes its
		execution.

		The defer statement basically pushes a function call to a list, and the list of
		saved calls is executed after the surrounding function returns. Defer is most
		commonly used to clean up resources, like closing a file handler
	*/
}

func main() {
	//panicTest(true)
	//fmt.Println("Hello")

	/*
		The panic caused the program to be terminated, which is why hello world was never printed. Instead, we got the panic message.

		So, now that we understand how panics work, an obvious question arises—how do we
		capture a panic and prevent it from killing our program?

		Before we answer that question, we first need to introduce the concept of defer.
		The defer keyword can be used to indicate that a piece of code must only be
		executed after the surrounding function returns.
	*/

	doSomething()
	/*
			prints out:
					0
					1
					2
					3
					4
					doSomething() just ended 2
					doSomething() just ended

		The defer statements typically enter a stack data structure, which means they
		execute based on the first-in-last-out rule. So, this basically means that the
		first defer statement in the code will execute last, while the next one will
		execute right before it and so on.

	*/

	panicTest2(true)
	fmt.Println("Hello!")
}

func checkPanic() {
	if r := recover(); r != nil {
		fmt.Println("A panic was captured, message:", r)
	}
}

func panicTest2(p bool) {
	// in here we use combination of defer and recover
	defer checkPanic()
	if p {
		panic("panic requested")
	}

	/*
		As you can see, we utilized a combination of defer and the recover() function to
		capture the panic to prevent it from terminating our program. If no panic
		occurred, the recover() function will return nil. Otherwise, the recover()
		function will return the error value of the panic. If we use recover() alone, it
		won't be effective without being combined with defer.
	*/
}
