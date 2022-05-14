package main

import "fmt"

// Interfaces

/*
	An interface can be very simply described as a Go type that hosts a collection of methods.
*/

// The preceding interface defines two methods—GetName() and GetAge().

type MyInterface interface {
	GetName() string
	GetAge() int
}

type Person struct {
	name string
	age  int
}

func (p Person) GetName() string {
	return p.name
}

func (p Person) GetAge() int {
	return p.age
}

/*
	In Go, an interface can be implemented by other types, like Go structs. When a
	Go type implements an interface, a value of the interface type can then hold
	that Go type data. We'll see what that means very shortly.

	A very special feature in Go is the fact that for a type to implement or inherit
	an interface, the type only needs to implement the methods of said interface.

	In other words, the Person struct type from the preceding piece of code
	implements the myInterface interface type. This is due to the fact that the
	Person type implements GetName() and GetAge(), which are the same methods that
	were defined by myInterface. So, what does it mean when Person implements
	MyInterface?
*/

type PersonWithTitle struct {
	name  string
	title string
	age   int
}

func (p PersonWithTitle) GetName() string {
	return p.title + " " + p.name
}

func (p PersonWithTitle) GetAge() int {
	return p.age
}

func main() {
	var myInterfaceValue MyInterface
	var p = Person{}
	p.name = "Jack"
	p.age = 39
	myInterfaceValue = p
	name := myInterfaceValue.GetName()
	age := myInterfaceValue.GetAge()

	fmt.Println("name:", name, "\nage:", age)

	p2 := Person{"Alice", 26}
	PrintNameAndAge(p2)

	p3 := PersonWithTitle{"Alice", "Dr", 26}
	PrintNameAndAge(p3)

	// There are cases where you might want to get back the concrete type value from
	// an interface value. Go includes a feature called type assertion that can be
	// used for just that. Here is the most useful form of type assertion:

	person, ok := myInterfaceValue.(Person)
	if ok {
		fmt.Println("person is a Person", person)
	}
}

func PrintNameAndAge(i MyInterface) {
	fmt.Println("name:", i.GetName(), "\nage:", i.GetAge())

}
