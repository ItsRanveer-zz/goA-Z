// go10.go

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

//Assigning println function to a variable so that we can use it instead
var p = fmt.Println

func main() {

	//Random Numbers
	p(rand.Intn(10)) //For generating random numbers from 0 to 10
	p(rand.Intn(10), ",", rand.Intn(50), ",", rand.Intn(100))
	p(rand.Int())               //For generating some Int random number
	p(rand.Float64())           //For generating some random float number of 64 bit which will be 0.0 <= f < 1.0
	p(rand.Float32())           //For generating some random float number of 32 bit
	p((rand.Float64() * 5) + 5) //For generating some random float number of 64 bit which will be 5.0 <= f < 10.0

	source := rand.NewSource(20) //To produce some deterministic random values we will give it our own seed. This will return a new random Source seeded with given value
	random := rand.New(source)   //Will return a Rand which uses Source to generate some random values
	p(random.Intn(20))           //Now you can call functions on random same like on rand
	p(random.Intn(100))
	p(random.Int())
	p(random.Float64())

	//Number Parsing
	f, _ := strconv.ParseFloat("2.45", 64) //Parsing Float Values, Here it will parse 64 bits of precision
	fmt.Println(f)
	f, _ = strconv.ParseFloat("2.45", 32) //Here it will parse 32 bits of precision
	fmt.Println(f)
	i, _ := strconv.ParseInt("256", 0, 64) //Parsing Int Values, 0 means infer the base from string, 64 is for result bits
	fmt.Println(i)
	d, _ := strconv.ParseInt("0x1b9", 0, 64) //Parsing Int Values will also parse hex-formatted numbers.
	fmt.Println(d)
	u, _ := strconv.ParseUint("598", 0, 64) //ParseUint is for Unsigned integers
	fmt.Println(u)
	k, _ := strconv.Atoi("265") //Atoi is for parsing int with base 10
	fmt.Println(k)

	//Exit, to exit a program with non-zero status we use os.Exit()
	//Even defer won't be executed after exit
	defer fmt.Println("I am defer") //Defer to print at the end of main

	fmt.Println("Before exit")
	connectDB()               //Call function which contains exit
	fmt.Println("After exit") //Will not be printed
}

func connectDB() {
	fmt.Println("Connecting DB...")
	//For a usecase if any problem occurred while connecting DB we can exit with some status
	//Exiting the program with status 4, you can define any status. After this nothing will be executed
	os.Exit(4)
}
