// go20.go

package main

import (
	"fmt"
	"os"
)

func main() {

	//Environment Variables, these let us set some config information
	//Setenv sets a env variable with key and value passed to it
	os.Setenv("firstkey", "firstvalue")
	fmt.Println("firstkey:", os.Getenv("firstkey"))     //Getenv gets the env var by passing the key
	fmt.Println("unknownkey:", os.Getenv("unknownkey")) //If key not set than Getenv will return empty string
	fmt.Println("Path :", os.Getenv("PATH"))            //PATH is one default env var set in the system
	fmt.Println()
	//Environ gives all the env variabels, there we will also get "firstkey" and "PATH" as one of env
	for _, envvar := range os.Environ() { //Printing env var one by one
		fmt.Println(envvar)
	}
}
