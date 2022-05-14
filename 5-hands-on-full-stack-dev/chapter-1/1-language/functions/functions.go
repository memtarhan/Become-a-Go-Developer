package main

import "fmt"

// Functions and Closures

// Here is what a function with arguments would look like:
func addIntegers(a int, b int) {
	total := a + b
	fmt.Println("", a, "+", b, "is", total)
}

// Since the a and b arguments from the preceding code are of the same type, we can also do this:
func addDecimals(a, b float64) {
	total := a + b
	fmt.Println("", a, "+", b, "is", total)
}

// Now, let's say we want to return a value from our function. Here is what this would look like:
func calculateSum(a, b int) int {
	total := a + b
	return total
}

// Go also allows multiple returns, so you can do this:
func addAndSubstract(a, b int) (int, int) {
	sum := a + b
	dif := a - b
	return sum, dif
}

// In Go, there is a concept known as named returns, which basically means that
// you can name your return values in the function header. Here is what this
// looks like:
func addSubtract(a, b int) (sum, dif int) {
	sum = a + b
	dif = a - b
	return
}

// Functions are also first-class citizens in the Go language, which means that
// you can assign a function to a variable and use it as a value. Here is an
// example:
var adder = func(a, b int) int {
	return a + b
}
var subtractor = func(a, b int) int {
	return a - b
}

// Because of this, you can also pass functions as arguments to other functions:
func execute(op func(int, int) int, a, b int) int {
	return op(a, b)
}

// Go also supports the concept of variadic functions. A variadic function is a
// function that can take an unspecified number of arguments. Here is an example
// of an adder function that takes an unspecified number of int arguments and
// then adds them:
func infiniteAdder(inputs ...int) (sum int) {
	for _, v := range inputs {
		sum += v
	}
	return
}

func main() {

	addIntegers(1, 2)
	addDecimals(3.14, 1.0)

	sum := calculateSum(5, 4)
	fmt.Println("Sum of", 5, "and", 4, "is", sum)

	add, dif := addAndSubstract(10, 5)
	fmt.Println("Sum of", 10, "and", 5, "is", add)
	fmt.Println("Dif of", 10, "and", 5, "is", dif)

	var addResult = adder(2, 3)
	fmt.Println("addResult", addResult)
	var subResult = subtractor(3, 2)
	fmt.Println("subResult", subResult)

	var value = execute(adder, 3, 2)
	fmt.Println("value", value)

	var infinite = infiniteAdder(1, 2, 2, 2) // 1 + 2 + 2 + 2
	fmt.Println("infinite", infinite)

}
