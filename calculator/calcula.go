package main

import "fmt"


func main(){
	var op string
	var number1 int
	var number2 int

	fmt.Println("Enter the operation either add, substraction, divide, multiply")
	fmt.Scan(&op)
	
	fmt.Println("Enter first number")
	fmt.Scan(&number1)
	fmt.Println("Enter Second number")
	fmt.Scan(&number2)

	switch {
		case op == "add":
			fmt.Println(number1 + number2)
		case op == "substraction":
			fmt.Println(number1 - number2)
		case op == "divide":
			fmt.Println(number1 / number2)
		case op == "multiply":
			fmt.Println(number1 * number2)
		default:
			fmt.Println("Not an operation")
	
	}


}