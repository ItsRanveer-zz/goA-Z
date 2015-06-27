// go5.go
package main

import (
	"errors"
	"fmt"
)

func main() {

	// Defer (to execute a function after the enclosing function finishes)
	// Often used to keep ur function call together but execute it at the end which may contain some cleanup activity.
	printOne()

	// Recover (Recover is to Catch if an ecxeption or error occur's and than continue execution)
	fmt.Println("Dividing by Zero : ", safeDiv(3, 0)) //This will create Divide By Zero Error
	fmt.Println(safeDiv(6, 2))                        //This will work fine

	// Panic (To create a runtime panic - exception)
	demPanic()
	fmt.Println("Hello I am after Panic")

	//Error
	if result, err := doSumCreateError(2, 2); err != nil { //Here we are also doing inline error check
		fmt.Println("Returned Sum: ", result, ", Error Occured:", err)
	}
}

//Defering a function
func printOne() {
	defer printThree() // This will be executed when printOne finishes
	defer printTwo()   // This will be also be executed when printOne finishes but before printThree
	fmt.Println("Print One")
}
func printTwo()   { fmt.Println("Print Two") }
func printThree() { fmt.Println("Print Three") }

//Recover an Error
func safeDiv(num1, num2 int) int {
	defer func() { // Making it defer so that it will be executed at the end even after error occurs
		recover() // To recover the error, we can also print the error. fmt.Println(recover())
	}()
	soln := num1 / num2
	return soln
}

//Panic
func demPanic() {
	defer func() { // Making it defer so that it will be executed at the end and recover the panic
		fmt.Println(recover())              //Recovering and printing panic
		fmt.Println("This will be printed") //Will be printed because inside defer func
	}()
	panic("I am a PANIC")                   //Creating Panic
	fmt.Println("This will not be printed") //Will not be printed because panic occurred
}

//Returning Error Or Throwing as in other languages
func doSumCreateError(num1, num2 int) (int, error) { //We can return error as a return type
	return num1 + num2, errors.New("I am the error Message") //Here we are returning error in our second return type and also assigning error Message
}
