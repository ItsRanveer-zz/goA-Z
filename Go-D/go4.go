//go4.go

package main

//Go Functions and Closures and Recursive Function
import "fmt"

//Defining main function (There can and should be only one main function otherwise error - "main redeclared" or "undefined main")
func main() {

	//Defining functions
	//First Example
	listNums := []float64{2, 8, 5, 9, 5, 2}
	fmt.Println("Sum of Numbers :", addThemUp(listNums))

	//Second Example - Multiple Return Type
	number1, number2 := next2Values(5)
	fmt.Println(number1, number2)
	//Even Println is also a functions which returns two values, the length of string what it printed and error if occurred
	length, err := fmt.Println("Hello")
	fmt.Println("Length: ", length)
	fmt.Println("Error: ", err)

	//Third Example - Variadic functions
	fmt.Println(addMultiple(1, 5, 8, 9, 8)) //If you want to send the values of a slice as a parameter you can send it like this "sliceName..."

	//Fourth Example - First Class Function (Functions are first class objects in Go)
	//We can store function in fields of a struct, pass them as arguments to other functions and use them as return values to other functions
	ownInstance := OwnType{func(first string) { fmt.Println("Hello,", first) }}                         //Creating an instance of struct
	returnedFunction := sayHello(ownInstance.newfunction, "By First Class Function Passed as Argument") //Calling a function by passing another function as argument, this function also returns new function
	returnedFunction("By First Class Function Got as returned value")                                   //Calling the function which was returned

	//Closure (or Anonymous Function). Used to define function inside function or assign function to a variable. Lexical Scoping.
	//First Example
	multipleNum := 3
	doubleNum := func() int { // Here we are assigning a function to doubleNum and using multipleNum variable.
		multipleNum *= 2 //doubling the value of multipleNum
		return multipleNum
	}
	fmt.Println("Calling doubleNum First Time - ", doubleNum())  // doubleNum will have access to multipleNum variable declared outside the function which was assigned to it.
	fmt.Println("Calling doubleNum Second Time - ", doubleNum()) // so everytime we call doubleNum it will changes multipleNum globally.
	fmt.Println(multipleNum)                                     //Value changed after calling doubleNum

	//Second Example
	var addNum func() //Defined a variable of type func()
	addNum = trpNum() //Here we are calling a function which in turn is returning a function which we will store in addNum()
	//The returned function is using variable "a" inside it. This variable will have its scope inside addNum()
	addNum() //Every time we call addNum "a" will be incremented
	addNum()

	//Third Example
	increment, get := Incremental(5)                 //This function returns two types
	fmt.Println("Value of num from get() - ", get()) //First calling the get and finding the current value in "num". You passed 5 so you should get 5.
	increment()                                      //Incrementing the value in num. It should increment it by 1, so 6
	increment()                                      //Once Again. It should increment it by 1, so 7
	fmt.Println("Value of num from get() - ", get()) //Again checking the value of "num". Current value should be 7

	//Recursive Function
	fmt.Println(factorial(5)) //!5 = 5*4*3*2*1

}

//Function taking slice in parameter, doing sum of elements and returning one value of float64 type
func addThemUp(numbers []float64) float64 {
	sum := 0.0
	for _, val := range numbers {
		sum += val
	}
	return sum // Explicit return in Go, not automatic.
}

//Function returning multiple values. Here in this case it is returning two values of int type
func next2Values(number int) (int, int) {
	return number + 1, number + 2
}

//Function receiving multiple values as an argument. Also called variadic functions. Here you can pass multiple number of integers comma seperated to this function.
func addMultiple(numbers ...int) int {
	finalValue := 0
	for _, value := range numbers { //range will iterate through all the arguments
		finalValue += value
	}
	return finalValue
}

//Defining own types(struct) which contains a field which is a funciton.
type OwnType struct {
	newfunction func(string) //Function as a field of struct
}

//Function taking other function as argument and returning new function.
func sayHello(firstClass func(string), helloString string) func(string) {
	firstClass(helloString)           //Calling that function which came as an argument
	return func(firstClass2 string) { //Returning a new function
		fmt.Println("Hello,", firstClass2)
	}
}

//This functin is returning a function
func trpNum() func() {
	a := 0                //local variable to trpNum
	increment := func() { //Here we are returning a funcion and inside it we are using variable "a"
		a += 1 //This variable will have its scope even after the trpNum() finishes because we are returning this function in which it is used
		fmt.Println("Value of a - ", a)
	}
	return increment // increment is of type func() which we are returning
}

//Defining function type so that we can use, instead of defining it thereitself. This makes code more readable
type incType func()
type getType func() int

//This function is taking an int as input and returning two func() types
//Here num is local to Incremental but because it is used in the returning functions, so its scope will still persist.
//Also both returned function is using the same "num" so both will have same value, if increment will change the value of "num" than get will also have same value.
func Incremental(num int) (incType, getType) {
	increment := func() {
		num++
		fmt.Println("Value of num from increment() - ", num)
	}
	get := func() int {
		return num
	}
	return increment, get
}

//Recursive Function
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
