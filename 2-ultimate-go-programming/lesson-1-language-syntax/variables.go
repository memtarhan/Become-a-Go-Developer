package main

import (
	"fmt"
)

func main() {
	// Declare variables that are set to their zero value
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T -> %v\n", a, a)
	fmt.Printf("var b int \t %T -> %v\n", b, b)
	fmt.Printf("var c int \t %T -> %v\n", c, c)
	fmt.Printf("var d int \t %T -> %v\n", d, d)
	/*
		Prints out:

			var a int 	 int -> 0
			var b int 	 string ->
			var c int 	 float64 -> 0
			var d int 	 bool -> false

	*/

	// Declaring variables
	var aa int = 0
	var bb string = "This is not a string"
	var cc float64 = 3.1415
	var dd bool = true

	fmt.Printf("var aa int \t %T -> %v\n", aa, aa)
	fmt.Printf("var bb int \t %T -> %v\n", bb, bb)
	fmt.Printf("var cc int \t %T -> %v\n", cc, cc)
	fmt.Printf("var dd int \t %T -> %v\n", dd, dd)


	aaa := 0
	bbb := "This is not a string"
	ccc := 3.1415
	ddd := true

	fmt.Printf("var aaa int \t %T -> %v\n", aaa, aaa)
	fmt.Printf("var bbb int \t %T -> %v\n", bbb, bbb)
	fmt.Printf("var ccc int \t %T -> %v\n", ccc, ccc)
	fmt.Printf("var ddd int \t %T -> %v\n", ddd, ddd)
}
