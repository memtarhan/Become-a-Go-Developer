package main

import (
	"errors"
	"fmt"
	"strconv"
)

func greet() string {
	return "Hello!"
}

func greetWithName(name string) string {
	return "Hello, " + name + "!"
}

func greetWithNameAndAge(name string, age int) (greeting string) {
	greeting = "Hello, my name is " + name + " and I am " + strconv.Itoa((age)) + " years old"
	return
}

func divide(a int, b int) (int, error) {
	if b == 0 { 
		return 0, errors.New("cannot divide by 0")	
	}

	return a / b, nil
}

func main() {
	fmt.Println(greet())
	fmt.Println(greetWithName("Mehmet"))
	fmt.Println(greetWithNameAndAge("Mehmet", 27))
	fmt.Println(divide(10, 2))
	fmt.Println(divide(5, 0))
}