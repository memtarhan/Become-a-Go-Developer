package main

import "fmt"

// declare packege-level variables of type int 
var d, e, f int  // default value is 0 

// declare package-level variables of type bool 
var x, y, z = true, false, true 

func main() {
	fmt.Println("d, e, f:", d, e, f)

	// override the default value 
	d = 1_000_000
	fmt.Printf("d: %d\n", d)


	x, y, z := 0, 1.25, "hello"
	fmt.Println("x, y, z:", x, y, z)

	printVars()
}

func printVars() {
	fmt.Println("x, y, z:", x, y, z)
}