// go6.go
package main

import (
	"fmt"
)

func main() {
	//Pointer (Passing the reference or the memory address instead of the actual value)
	x := 10
	changePointerVal(&x)                      //Passing the pointer or memory address of x i.e. "&x"
	fmt.Println("x = ", x)                    //Value of x changed after calling the changePointerVal function
	fmt.Println("Memory address of x = ", &x) //Printing the memory address of x

	yPtr := new(int)                                                       //Can also create a pointer with new() by default value 0
	fmt.Println("Memory Address of yPtr Pointer = ", &yPtr)                //yPtr is of type *int (pointer) but that will also have an address in the memory
	fmt.Printf("Memory Address of yPtr Pointer = %p\n", &yPtr)             //Can use %p for string formatting
	fmt.Println("Value of yPtr which itself is an address = ", yPtr)       //yPtr is a pointer so its value will be an address
	fmt.Println("Value at that address which is stored in yPtr = ", *yPtr) //Value to which pointer is pointing to. We have just created and not initialized so default value = 0
	changePointerVal(yPtr)                                                 //Once again calling changePointerVal to change the value to which yPtr points to.
	fmt.Println("Value changed at that address which is stored in yPtr = ", *yPtr)
}

//Function which accepts pointer
func changePointerVal(a *int) { // a is a pointer of type int which will contain only memory address
	*a = 20 // Here we dereference and assign the value to that pointer, this will put the value at that memory address. So the actual value will change
}
