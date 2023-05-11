package data

import "fmt"

type Duration float32 // in hours

type Course struct {
	Id int
	Name string
	Slug string
	Legacy bool
	Duration Duration
	Instructor Instructor
}


func (c Course) String() string {
	return fmt.Sprintf("%s -- %s", c.Name, c.Instructor.FirstName)
}