package main

import "fmt"

// Pointers

func main() {

	/*
		The concept of pointers is simple—a pointer is a language type that represents
		the memory locations of your values. Pointers in Go are used everywhere, and
		that's because they give the programmer a lot of power over the code. For
		example, having access to your value in memory allows you to change the original
		value from different parts of your code without the need to copy your value
		around.
	*/

	/*
		In Go, to create a pointer, you just append * in front of the data type of your
		value. For example, here is a pointer to an int value:
	*/
	var iptr *int
	fmt.Println("Initial value of iptr is ", iptr)

	/*
		As we mentioned in the previous section, the zero value of a pointer is nil. The
		behavior of nil is similar to null in languages like Java, that is, if you try
		to use a nil pointer, an error will get thrown.
	*/
	var x int = 5
	fmt.Println("Initial value of x is ", x)

	// We also want a pointer to point to the address of x for later use:
	var xptr = &x
	fmt.Println("Initial value of xptr is ", &xptr)

	/*
		The & operand here means that we want the address of x. Whenever you append the
		& operand before a variable, it basically means that we want to the address of
		that variable.
	*/

	/*
		What if we have a pointer, and we want to retrieve the value that it points to?
		This operation is called de-referencing, and here is how we can do it:
	*/
	y := *xptr
	fmt.Println("Initial value of y is ", y)

	/*
		What if we want to change the value that the pointer points to? We can still use
		de-referencing for that, and here is what this would look like:
	*/
	*xptr = 4
	fmt.Println("Later value of xptr is ", &xptr)
}
