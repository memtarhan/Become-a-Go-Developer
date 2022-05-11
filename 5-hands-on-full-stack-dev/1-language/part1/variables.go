package main

import "fmt"

// Variables and data types

// Declaring variables
var sx string
var s1, s2, s3 string

/*
	To initialize a variable with an initial value, Go offers a number of options.
	One option is to initialize the variable while also specifying the variable
	type. Here is what this looks like:
*/
var a1, a2, a3 string = "first", "second", "third"

/*
	Another option is to initialize the variables without specifying the data type.
	Here is what this would look like:
*/
var b1, b2, b3 = "first", "second", "third"

/*
	We can then mix data types with the following syntax:
*/
var s, i, f = "string", 10, 3.14

/*
	A popular way to declare and initialize multiple variables at once is as follows:
*/
var (
	s5 = "string"
	i5 = 10
	f5 = 3.14
)

func main() {
	/*
		If you are declaring and initializing your variable inside a function, you don't
		need to even use the var keyword. Instead, you can use :=. This is called type
		inference, since you infer the variable type from the provided value. Here is
		how we would declare and initialize the s, i, and f variables with type
		inference:
	*/
	s := "string"
	i := 10
	f := 3.14
	/*
		The var keyword, however, gives you more control since it allows you to
		explicitly specify the data type you would like to use for your variable.
	*/

	fmt.Println("string ", s, "integer ", i, "float ", f)

	/*
			Variables that are declared without an explicit initial values get assigned what
			is known as zero values. Here is a table for zero values:

		Numeric types	0
		Boolean types	false
		String type	""
		Pointers	nil

	*/

}
