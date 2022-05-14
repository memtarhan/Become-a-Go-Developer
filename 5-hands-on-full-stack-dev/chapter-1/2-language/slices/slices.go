package main

import "fmt"

// Slices
/*
	There is a very obvious limitation in Go's array data structure—you must specify
	the size whenever you declare a new array. In real life, there are numerous
	scenarios where we will not know the number of elements to expect beforehand.
	Almost every modern programming language comes with its own data structure to
	address this requirement. In Go, this special data structure is called a slice.
*/

func main() {
	// From a practical point of view, you can think of slices as simply dynamic
	// arrays. From a syntax point of view, slices look very similar to arrays,
	// except that you don't need to specify the size
	var mySlice []int
	mySlice = []int{1, 2, 3, 4, 5}
	fmt.Println("mySlice:", mySlice)

	var mySlice2 = []int{1, 2, 3, 4, 5}
	fmt.Println("mySlice2:", mySlice2)

	// Since slices can grow in size, we are also allowed to initialize an empty slice:
	var mySlice3 = []int{}
	fmt.Println("mySlice3:", mySlice3)

	// If you would like to set an initial number of elements in your slice without
	// having to write the initial values by hand, you can utilize a built-in
	// function called make:
	var mySlice4 = make([]int, 5)
	fmt.Println("mySlice4:", mySlice4)

	/*
		A slice can simply be considered as a pointer to a part of an array. A slice
		holds three main pieces of information:

		* A pointer to the first element of the subarray that the slice points to.
		* The length of the subarray that's exposed to the slice.
		* The capacity, which is the remaining number of items available in the original array.
		  The capacity is always either equal to the length or greater.
	*/

	fmt.Println("*********")

	var newSlice = []int{1, 2, 3, 4, 5}
	fmt.Println("newSlice:", newSlice)
	var subSlice = newSlice[2:4]
	capacity := cap(subSlice)
	fmt.Println("Capacity of", subSlice, "is", capacity)
	length := len(subSlice)
	fmt.Println("Length of", subSlice, "is", length)

	/*
		You might be wondering by now, why should I care about the differences between
		length and capacity? I can just use the length and ignore the capacity
		altogether, since the capacity only gives you information about a hidden
		internal array.

		The answer is very simple—memory utilization. What if mySlice had 100,000
		elements instead of just five? This means that the internal array would have had
		100,000 elements as well. This huge internal array will exist in our program's
		memory as long as we use any sub-slices extracted from mySlice, even if the
		sub-slices we use only care about two elements.

		To avoid that kind of memory bloat, we need to explicitly copy the fewer
		elements we care about into a new slice. By doing this, once we stop using the
		original large slice, Go's garbage collector will realize that the huge array is
		not needed anymore and will clean it up.

		So, how do we achieve that? This can be done through a built-in function called copy:
	*/

	// let's assume this is a huge slice
	var myBigSlice = []int{1, 2, 3, 4, 5, 6}
	// now here is a new slice that is smaller
	var mySubSlice = make([]int, 2)
	// we copy two elements from myBigSlice to mySubSlice
	copy(mySubSlice, myBigSlice[2:4])
	fmt.Println("mySubSlice:", mySubSlice)

	fmt.Println("*********")

	/*
		We keep saying that slices are like dynamic arrays, but we haven't seen how to
		actually grow the slice yet. Go offers a simple built-in function called append,
		which is used to add values to a slice. If you reach the end of your slice
		capacity, the append function will create a new slice with a bigger internal
		array to hold your expanding data. append is a variadic function, so it can take
		any number of arguments. Here is what this looks like:
	*/

	var appendingSlice = []int{1, 2}
	fmt.Println("appendingSlice:", appendingSlice)
	appendingSlice = append(appendingSlice, 3, 4, 5, 6)
	fmt.Println("appendingSlice:", appendingSlice)

	//Initialize a slice with length of 3
	var mySlice5 = make([]int, 3)
	fmt.Println("mySlice5:", mySlice5)
	//initialize a slice with length of 3 and capacity of 5
	var mySlice6 = make([]int, 3, 5)
	fmt.Println("mySlice6:", mySlice6)
	//Initialize a slice with length of 3, and capacity of 3
	var mySlice7 = make([]int, 3)
	fmt.Println("mySlice7:", mySlice7)
}
