package main

import "fmt"

// Methods

/*
	A method is basically a function that you attach to a type.
	For example, let's assume we have a struct type called Person:
*/

type Person struct {
	name string
	age  int
}

// Go allows us to attach a method to that type like this:

// GetName returns name of a person
func (p Person) GetName() string {
	return p.name
}

// The part between the func keyword and the function name, GetName(), is known as the receiver of the method.

// GetAge return age of a person
func (p Person) GetAge() int {
	return p.age
}

// Type Embedding\

/*
	But what if you would like a struct to inherit the methods of another struct?
	The closest feature that the Go language offers to the concept of inheritance is
	known as type embedding
*/

type Student struct {
	Person
	studentId int
}

func (s Student) GetStudentId() int {
	return s.studentId
}

/*
	Notice that in the preceding code, we included the type Person inside the struct
	definition of type Student, without specifying a field name. This will
	effectively make the Student type inherit all the exported methods and fields of
	the Person struct type. In other words, we can access the methods and fields of
	Person directly from an object of type Student:
*/

func main() {
	var p = Person{
		name: "Jason",
		age:  29,
	}

	var name = p.GetName()
	fmt.Println("Name:", name)

	s := Student{}
	s.Person = p
	s.studentId = 1

	fmt.Println("Student's name:", s.GetName())
	fmt.Println("Student's age:", s.GetAge())

	// In Go, when a type gets embedded inside another type, the exported methods and
	// fields of the embedded type are said to be promoted to the parent or embedding type.
}
