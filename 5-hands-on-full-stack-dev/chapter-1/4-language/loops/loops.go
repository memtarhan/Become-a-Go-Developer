package main

import "fmt"

// Loops

/*
	In Go, there is a single keyword that you can use when you want to write a loop—for.
	There are no other keywords to indicate a loop in Go.
*/

func main() {
	/*
	*	An initial value (i:=1) in your code—this is optional
	*	A condition to indicate whether to keep iterating or not (i<=10)
	*	The value of the next iteration (i++)
	 */
	for i := 0; i < 3; i++ {
		fmt.Println("i:", i)
	}

	/*
		What if we have a slice or an array and we want to iterate over it in a loop? Go
		comes to the rescue with the concept of for .. range. Let's assume we have a
		slice called myslice and that we want to iterate over it. Here is what the code
		would look like:
	*/

	mySlice := []string{"one", "two", "three", "four"}
	for i, item := range mySlice {
		fmt.Println("i:", i, "item:", item)
	}

	// What about if we only care about the index? For that, we can do this:
	for i := range mySlice {
		fmt.Println("i:", i)
	}

	for i := range mySlice {
		mySlice[i] = "other"
	}

	fmt.Println(mySlice)

	/*
		Go supports the break and continue keywords. The break keyword inside a loop
		would cause the loop to break, even if it is not done. The continue keyword, on
		the other hand, will force a loop to jump to the next iteration.
	*/
}
