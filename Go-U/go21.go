// go21.go

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	fmt.Println("Spawning")
	//Spawning, in go one process can load and execute other processes,
	//which can be a go or non-go process.

	//To execute a external process on cmd we use Command, we pass our process name which
	//we want to execute and it will return a cmd object that needs to be executed on the command prompt.
	//Here date process is an external process that will be represented by returned object
	dateCmd := exec.Command("date")
	//Output will run that command and return its standard output in bytes and also err if any
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Print("Date: ", string(dateOut)) //Printing the std output returned form cmd

	//Second Example
	//You can also pass arguments to oyur process, here process is echo and args is string
	echoCmd := exec.Command("echo", "Hello, I m an argument to echo")
	//Output will run that command and return its standard output in bytes and also err if any
	echoOut, err := echoCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Print("Echo: ", string(echoOut))

	//Third Example
	//Here we will hold pipes to the standard input, output and error and we will pass values to our process
	//First we will create a process which will print lines maching a pattern, here "hello"
	grepCmd := exec.Command("grep", "hello")
	//Now we will create pipes to our process
	grepIn, _ := grepCmd.StdinPipe()   // StdinPipe returns a pipe that will be connected to the command's standard input when the command starts.
	grepOut, _ := grepCmd.StdoutPipe() // StdoutPipe returns a pipe that will be connected to the command's standard output when the command starts.
	//You can also connect a pipe to std err to get the errors using StderrPipe
	//grepErr, _ := grepCmd.StderrPipe() //StderrPipe returns a pipe that will be connected to the command's standard error when the command starts.

	grepCmd.Start() //Start the process now using Start

	grepIn.Write([]byte("hello grep\ngoodbye grep")) //Write to the standard input using the input pipe, here we are piping two lines to process
	grepIn.Close()                                   //Close the input pipe
	grepBytes, _ := ioutil.ReadAll(grepOut)          //ReadAll will read all the data which came in the output pipe
	grepCmd.Wait()                                   //Wait will wait for the command to exit and release all resources

	fmt.Print("Grep hello: ", string(grepBytes)) //Only first line contains hello so that will be printed

	//For spawning full command and arguments as a string we can use bash with -c option
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h") //Last string will be our command with options and arguments
	lsOut, _ := lsCmd.Output()
	fmt.Println("List directory contents ls -a -l -h: ", string(lsOut))

	fmt.Println("Exec'ing")
	//Exec'ing Processes, to completely replace the current go process with another non-go one
	//To execute ls command we need to give its path in the system
	binary, _ := exec.LookPath("ls") //LookPath will return the path to the binary or err if occurred

	args := []string{"ls", "-a", "-l", "-h"} //syscall.Exec reads a slice so we put args in slice

	env := os.Environ() //syscall.Exec also uses environment variables

	//After calling Exec our program will be replaced by ls program
	execErr := syscall.Exec(binary, args, env) //Pass path to binary, arguments and env variables to Exec
	if execErr != nil {                        //Ecec returns err if occurred
		panic(execErr)
	}

	//Once our program is replaced by ls without any errors then it will not do the remaining things
	//So this hello will not be printed
	fmt.Println("Hello")
}
