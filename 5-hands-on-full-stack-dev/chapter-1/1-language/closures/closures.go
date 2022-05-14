package main

// Closures

/*
	A function can also be a closure. A closure is a function value that's bound to
	variables outside its body. This means that a closure can access and change
	values on those variables
*/

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

/*
	The closure in the preceding example has access to the sum variable, which means
	that it will remember the current value of the sum variable, and will also be
	able to change the value of the sum variable
*/

func main() {
	// when we call "adder()",it returns the closure
	sumClosure := adder() // the value of the sum variable is 0
	sumClosure(1)         // now the value of the sum variable is 0+1 = 1
	sumClosure(2)         // now the value of the sum variable is 1+2=3
}
