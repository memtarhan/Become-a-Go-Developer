package main

import "fmt"

// The switch statement

func getNumber() int { return 10 }

func main() {
	number := getNumber()

	switch number {
	case 5:
		fmt.Println(5)
	case 6:
		fmt.Println(6)
	default:
		fmt.Println("default case")
	}

	// If you haven't noticed already, there is no break keyword.
	// In Go, each case breaks automatically, and doesn't need to be told to do so.

	switch number := getNumber(); number {
	case 5:
		fmt.Println(5)
	case 6:
		fmt.Println(6)
	default:
		fmt.Println("default case")
	}

	// In Go, a switch statement can act like a group of if else.
	// This gives you the ability to write long if else chains with much nicer code:

	switch {
	case number == 5:
	case number > 10:
	default:
	}

	// In some scenarios, you want your switch cases not to break automatically,
	// and instead fall through to the next case.
	// For this, you can use the fallthrough keyword:

	switch {
	case number == 5:
		fallthrough
	case number > 10:
	default:
	}
}
