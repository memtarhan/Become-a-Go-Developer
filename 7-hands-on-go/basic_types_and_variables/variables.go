package main

import (
	"fmt"
	"unicode"
)

// declare packege-level variables of type int
var d, e, f int  // default value is 0 

// declare package-level variables of type bool 
var x, y, z = true, false, true 

// declare a constant representing pi
const pi = 3.14158

const a rune = 'a'

func main() {
	fmt.Println("d, e, f:", d, e, f)

	// override the default value 
	d = 1_000_000
	fmt.Printf("d: %d\n", d)


	x, y, z := 0, 1.25, "hello"
	fmt.Println("x, y, z:", x, y, z)

	printVars()

	fmt.Printf("pi: %v - %T\n", pi, pi)
	fmt.Printf("a: %c - %T\n", a, a)

	// use unicode package to confirm that the rune is a letter
	unicode.IsLetter(a)

	// declare float and convert it to an unsigned int 
	a := 42.5 
	b := uint(a)
	fmt.Printf("%v=%T, %v=%T", a, a, b, b)
	fmt.Println()

	doPointers()
}

func printVars() {
	fmt.Println("x, y, z:", x, y, z)
}

func doPointers() {
	// create a variable of type *T where T is an int 
	var a *int 

	// declare and assing 'b' variable of type int 
	b := 100 

	// assign the address of b to a 
	a = &b

	// print out the value of a which is the address of b 
	fmt.Println(a)

	// print out the value at the address of b 
	fmt.Println(*a)

}