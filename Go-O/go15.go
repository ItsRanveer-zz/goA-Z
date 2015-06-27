// go15.go
// For running this program you have to build it first
// $ go build gocommandlineargs.go
// Then you have to run your program and pass 4 arguments
// ./go15 a b c d
package main

import (
	"fmt"
	"os"
)

func main() {
	p := fmt.Println

	//For accessing the arguments passed to your program while
	//running from command line we use os.Args
	argsWithPath := os.Args //Args will give you a slice which will contain the
	//path to your program at 0 position than the number of args passed in the rest
	p(argsWithPath)

	argsWithoutPath := os.Args[1:] //To get only args we can slice it like [1:]
	p(argsWithoutPath)

	arg := os.Args[3] //We can also access the arg through its index
	p(arg)
}
