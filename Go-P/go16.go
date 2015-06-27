//go16.go
//For running first build the program
//$ go build go16.go
//Now run the binary passing diffrent flags and args and see the result
//$ ./go16 -word=opt -numb=7 -fork -svar=flag
//$ ./go16 -word=opt
//$ ./go16 -word=opt a1 a2 a3
//$ ./go16 -word=opt a1 a2 a3 -numb=7
//$ ./go16 -h
//$ ./go16 -wat

package main

import (
	"flag"
	"fmt"
)

func main() {
	//For accessing and parsing command line flags we use flag package
	//Defining string flag we use String, with flag name, default value and description as args
	wordPtr := flag.String("word", "foo", "a string") //String will return a string pointer

	//Defining int flag we use Int, with flag name, default value and description as args
	numbPtr := flag.Int("numb", 42, "an int") //Int will return a int pointer

	//Defining bool flag we use Bool, with flag name, default value and description as args
	boolPtr := flag.Bool("fork", false, "a bool") //Bool will return a bool pointer

	//Defining string flag using a predefined variable in program using StringVar
	var svar string //This is our predefined variable
	//Here first arg will be the pointer to var in which we have to store the value of the flag
	flag.StringVar(&svar, "svar", "bar", "a string var")

	//Now after defining the flags and before acessing them we have to parse command line args
	//Parse will executue the command line parsing
	//It will parse all the command line flags from os.Args[1:]
	flag.Parse()

	//Now we have parsed our command line flags and args we can use them
	fmt.Println("word:", *wordPtr) //Because our flags are pointers, so we need to dereference them to get actual value
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar) //svar is variable and when we created a flag of it we gave the pointer so
	//that pointer will atomatically store the value in that variable
	fmt.Println("tail:", flag.Args()) // Apart from flags if any trailing positional arguments are there than we can get it using Args
}
