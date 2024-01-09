package main

import "fmt"

// Structs

/*
	A struct in Go is a data structure that is composed of fields, where each field has a type.

	If the Go struct field names start with lower case letters, they will not be
	visible to external packages. If you want your struct or its fields to be
	visible to other packages, start the struct and/or field name with upper case
	letters.
*/

type myStruct struct {
	intField    int
	stringField string
	sliceField  []int
}

/*
	The preceding code creates a struct type that is called myStruct, which contains three fields:

	intField of type int
	stringField of type string
	sliceField of type []int
*/

func main() {
	var s = myStruct{
		intField:    3,
		stringField: "three",
		sliceField:  []int{1, 2, 3},
	}

	fmt.Println("s:", s)

	// or
	var s2 = myStruct{4, "four", []int{1, 2, 3, 4}}
	fmt.Println("s2:", s2)

	var s3 = myStruct{}
	s3.intField = 1
	s3.stringField = "one"
	s3.sliceField = []int{1}
	fmt.Println("s3:", s3)

	// You can obtain a pointer to a struct by doing this:
	var sPtr = &myStruct{
		intField:    2,
		stringField: "two",
		sliceField:  []int{1, 2},
	}
	fmt.Println("sPtr:", sPtr)

	// A dot notation can be used with a Go struct pointer, since Go will understand
	// what needs to be done without the need to do any pointer de-referencing:
	var s5 = &myStruct{}
	s5.intField = 5
	s5.stringField = "five"
	s5.sliceField = []int{1, 2, 3, 4, 5}
}
