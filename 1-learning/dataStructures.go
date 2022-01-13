package main

import (
	"log"
	"time"
)


type User struct {
	FirstName string
	BirthDate time.Time
}

// Pointing to User
func (m *User) getFirstName() string {
	return m.FirstName
}

func (m *User) getAge() int {
	difference := time.Now().Sub(m.BirthDate)
	age := difference.Hours() / (365 * 24)
	return int(age)
}

func main() {
	var myVar User
	myVar.FirstName = "Mehmet"

	birthdate := time.Date(2000, time.Month(2), 21, 1, 10, 30, 0, time.UTC)

	myVar2 := User {
		FirstName: "Mem",
		BirthDate: birthdate,
	}

	log.Println("myVar is set to", myVar.getFirstName())
	log.Println("myVar2 is set to", myVar2.getFirstName())

	log.Println("", myVar2.FirstName, "'s age is", myVar2.getAge())
}
