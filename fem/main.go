package main

import (
	"fmt"
	"goworkspace/go/server/data"
	"time"
)

func main(){
	var just data.Instructor
	mclean := data.Instructor {} // new instructor
	just.FirstName = "Justice"
	just.LastName = "McLean"
	just.Score = 98
	just.Id = 1

	fmt.Println(just.Print())

	mclean.FirstName = "McLean"
	mclean.LastName = "Duodu"
	mclean.Score = 99
	mclean.Id = 2
	fmt.Println(mclean.Print())

	mam := data.Instructor{Id: 3, FirstName: "Mamle", LastName: "Esther", Score: 97}
	fmt.Println(mam.Print())

	kyle := data.NewInstructor("Kyle", "Simpson")
	fmt.Println(kyle.Print())


	goCourse := data.Course{Name: "Go Fundamentals For Web Developers", Instructor: just, Id: 1, Duration: 90}

	fmt.Println(goCourse.String())

	swiftWS := data.WorkShop{
		Course: data.Course{Name: "Swift for iOS", Instructor: mam},
		Date: time.Now().Local(),
	}

	fmt.Printf("%v",swiftWS)
}