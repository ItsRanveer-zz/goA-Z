// go14.go
// Run this file using below command in curret directory
// $ cat text.txt | go run go14.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//Line filter are programs which are used to read inputs from stdin,
	//process it and print some derived result to stdout
	//Here we are scanning the input line by line and making it to Upper case and printing out on Stdout
	scanner := bufio.NewScanner(os.Stdin) //Buffered scanner to scan the unbuffered Stdin

	//Scan will scan the buffer token by token or line by line and returns false if reached the end or if err
	for scanner.Scan() {
		upperCase := strings.ToUpper(scanner.Text()) //Text returns the token or a line as a string
		fmt.Println(upperCase)
	}

	//If any error occurred while Scan than we can get that using Err on scanner
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err) //Fprintln will write the err on the Stderr
		os.Exit(1)                             //Exiting after the error
	}
}
