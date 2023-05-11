package data

import "fmt"

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

// factory

func NewInstructor(firstname string, lastname string) Instructor {
	return Instructor{FirstName: firstname, LastName: lastname}
}

func (i Instructor) Print() string {
	return fmt.Sprintf("Instructor: %s, %v Score: %d", i.FirstName, i.LastName, i.Score)
}
