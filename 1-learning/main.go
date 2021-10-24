package main

import "fmt"

func main() {
	fmt.Println("Here is Go")

	var whatToSay string
	whatToSay = "Goodbye, cruel world"

	var i int
	i = 2021

	fmt.Println(whatToSay)
	fmt.Println("i is set to:", i)

	whatWasSaid := saySomething()

	fmt.Println("The function returned:", whatWasSaid)

	name, age, country := introduceYourself()
	var ageString = fmt.Sprintf("%d", age)

	var introduction = "My name is " + name + " I am " + ageString + " years old and live in " + country
	fmt.Println(introduction)
}

func saySomething() string {
	return "Something"
}

func introduceYourself() (name string, age int, country string) {
	var mName = "Mehmet"
	var mAge = 2021 - 1995 
	var mCountry = "Portugal"

	return mName, mAge, mCountry
}
