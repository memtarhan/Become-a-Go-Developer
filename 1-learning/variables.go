package main

import (
	"log"
	"time"
)

/// Scopes
// var s string // Declare
var s = "seven" // Declare & Initialize

/// Structs

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
	isYoung     bool

	/*
		Ok, what's with the capitals?
			In Go, if we want to have a variable available outside of the current package(Here the package is main)
			we capitalize it. Otherwise it would be private to current package.
			In above example:
			Age is a public variable where isYoung is private
	*/
}

var birthDate time.Time

func main() {
	// fmt.Println("Here is Go")

	// var whatToSay string
	// whatToSay = "Goodbye, cruel world"

	// var i int
	// i = 2021

	// fmt.Println(whatToSay)
	// fmt.Println("i is set to:", i)

	// whatWasSaid := saySomething()

	// fmt.Println("The function returned:", whatWasSaid)

	// name, age, country := introduceYourself()
	// var ageString = fmt.Sprintf("%d", age)

	// var introduction = "My name is " + name + " I am " + ageString + " years old and live in " + country
	// fmt.Println(introduction)

	// var s2 = "six"
	// s := "eight"

	// log.Println("s is", s)
	// log.Println("s2 is", s2)

	// saySomething(s2)

	user := User{
		FirstName: "Mehmet",
		LastName:  "Tarhan",
	}

	log.Println(user.BirthDate)

}

func saySomething(s3 string) (string, string) {
	log.Println("s from the saySomething func is", s)
	return s3, "Something"
}

func introduceYourself() (name string, age int, country string) {
	var mName = "Mehmet"
	var mAge = 2021 - 1995
	var mCountry = "Portugal"

	return mName, mAge, mCountry
}
