package main

import "fmt"

// Arrays

/*
	An array is a common data structure that exists in any programming language. In
	Go, an array is a collection of values with the same data type, and a
	pre-defined size.
*/

func main() {
	var myArray [3]int // The preceding array is of type int and of size 3.
	// We can then initialize the array like this:
	myArray = [3]int{1, 2, 3}

	fmt.Println("myArray:", myArray)

	// Or, we can do this:

	myArray[0] = 4
	myArray[1] = 5
	myArray[2] = 6

	fmt.Println("myArray:", myArray)

	// Or, this:
	var myArray2 = [3]int{1, 2, 3}
	fmt.Println("myArray2:", myArray2)

	// Or, if we are declaring and initializing the array inside a function, we can use the := notation:
	myArray3 := [3]int{4, 5, 6}
	fmt.Println("myArray3:", myArray3)

	n := len(myArray)
	fmt.Println("Size of myArray:", n)

	/*
		Go also allows you to capture subarrays of your main array. To do that, you need
		to follow the following syntax:

		array[<index1>:<index2>+1]
	*/

	newArray := [5]int{1, 2, 3, 4, 5}
	// I can obtain a subarray from index two of my array till index three using the following syntax:
	var slice = newArray[2:4]
	fmt.Println("Slice at [2:4]:", slice)

	var startFromLeft = newArray[:2]
	fmt.Println("Slice at [:2]:", startFromLeft)
	var startFromRight = newArray[2:]
	fmt.Println("Slice at [2:]:", startFromRight)

	/*
		Let's say you do something like this:

		mySubArray := myarray[2:4]
		mySubArray is not merely a copy of a subpart of myarray.
		In fact, both arrays will point to the same memory
	*/

	mySubArray := newArray[2:4]
	mySubArray[0] = 2
	fmt.Println("newArray:", newArray) // [1 2 2 4 5]
}
