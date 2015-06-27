// go25.go

//Execute this program in command line
//$ go run go25.go
//and then interrupt using Ctrl+C wich will send SIGINT signal to our program
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//Signals, In go we can also send and receive signals using channels

	//We have to send os.Signal type values over the channels
	sigs := make(chan os.Signal, 1) //Here we create channel of os.Signal type

	//We also have to create a channel to receive notifications that we got our signal or not
	done := make(chan bool, 1)

	//Notify will register sigs channel to receive specified signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//Now we make a go routine call where we will hold to receive signal
	go func() {
		sig := <-sigs                           //Here sigs waits to receive signals of any specified type
		fmt.Println("\nSignal Received: ", sig) //When it receives the signal it will execute further and print it out
		done <- true                            //After we receive our singal we send a notification over done channel
	}()

	fmt.Println("Waiting for signal")

	<-done //Here done will wait untill it gets the notification from our goroutine above

	fmt.Println("Exiting Now")
}
