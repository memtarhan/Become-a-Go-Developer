package main

import "fmt"

// Maps

/*
	HashMaps are very popular, and extremely important data structures in any
	programming language. A map is a collection of key value pairs, where you use
	the key to obtain the value that corresponds to it. Using maps greatly speeds up
	your software due to the fact that with a map, retrieving a value through a key
	is a very quick operation.
*/
func main() {
	// The preceding code declares a map where the keys are of type int, and the values are of type string.
	var myMap map[int]string
	myMap = make(map[int]string)
	fmt.Println("myMap:", myMap)

	// You can't use a map before you initialize it, otherwise an error will be thrown.
	myMap2 := map[int]string{}
	fmt.Println("myMap2:", myMap2)

	// What if you want to initialize the map with some values?
	myMap3 := map[int]string{1: "first", 2: "second", 3: "third"}
	fmt.Println("myMap3:", myMap3)
	// Adding values
	myMap3[4] = "fourth"
	fmt.Println("myMap3:", myMap3)
	// Getting values
	var x = myMap3[4]
	fmt.Println("myMap3[4]:", x)

	// If the key 5 is not in "myMap", then "ok" will be false
	// Otherwise, "ok" will be true, and "y" will be the value
	y, ok := myMap3[5]
	if ok {
		fmt.Println("myMap3[4]:", y)
	}

	// You can delete a value from a map by using the built-in delete function:
	delete(myMap3, 4)

	fmt.Println("myMap3:", myMap3)
}
