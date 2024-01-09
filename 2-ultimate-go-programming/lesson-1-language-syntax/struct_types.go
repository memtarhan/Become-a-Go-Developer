package main

import (
	"fmt"
)

// Name type
type bill struct {
	flag bool
	counter int16 
	pi float32
}

// Name type
type alice struct {
	flag bool
	counter int16 
	pi float32
}

func main() {

	var e1 struct {
		flag bool
		counter int16 
		pi float32
	}

	fmt.Printf(("%+v\n"), e1)

	/// if only needs to used in one place
	// Literal type
	e2 := struct {
		flag bool
		counter int16 
		pi float32
	} {
		flag: true,
		counter: 10,
		pi: 3.141592,
	}

	fmt.Printf("%+v\n", e2)

	var b bill
	var a alice

	/// Go does not have implicit conversion
	// b = a // Won't work
	/// Explicit conversion
	b = bill(a)
	fmt.Println(b, a)

	/// But Go will allow literal type to name type implicit conversion. i.e
	b = e2 
	fmt.Println(b)

	/*
	 If we're dealing with values of a name type, there is no implicit conversion. 
	 You've gotta use your conversion syntax, you've gotta show your intent. 
	 When there's, when we're dealing with that literal type value then 
	 we're gonna get a little bit more flexibility
	*/
}