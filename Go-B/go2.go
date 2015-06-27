//go2.go

package main

import (
	"fmt"
)

func main() {
	// Logical operators in Go
	fmt.Println("true && false = ", true && false) // AND
	fmt.Println("true || false = ", true || false) // OR
	fmt.Println("!true = ", !true)                 // NOT

	//Relational Operators in Go
	// ==, !=, <, >, <=, >=

	//For Loops in Go
	//for loop with only a condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++ // i=i+1
	}

	//for loop with initialization;condition;increment/decrement
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	//for loop without conditon. We have to use break or return to come out of this loop otherwise it will be infinite loop
	for {
		fmt.Println("Inside For Loop")
		break
	}

	//for loop with range, range will iterates over each and every character of string and return unicode of that character
	for index, value := range "abc" {
		fmt.Println(index, value) //Will print Index of character and unicode of that
	}

	//If-Else Statement
	yourAge := 19
	if yourAge <= 16 {
		fmt.Println("You can't drive")
	} else if yourAge >= 18 {
		fmt.Println("You can Vote and drive")
	} else {
		fmt.Println("You can drive")
	}

	//if with declaration, you can declare any variable and use it in the scope of whole loop
	if a := 5; a < 0 {
		fmt.Println("Number is a negative number")
	} else if a < 10 {
		fmt.Println("Number is a positive number but less than 10")
	} else {
		fmt.Println("Number is a negative number")
	}

	//Switch Statement, default case is optional, use ',' to give multiple condition
	switch yourAge {
	case 16:
		fmt.Println("Go Drive")
	case 18, 19, 24:
		fmt.Println("Go Vote")
	default:
		fmt.Println("Sit at Home")
	}

	//switch without an expression, it can be used as alternative of if-else statement
	switch {
	case yourAge < 21:
		fmt.Println("You are not eligible for marrige")
	default:
		fmt.Println("Eligible for marrige")
	}
}
