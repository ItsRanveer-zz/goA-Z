// go22.go

package main

import (
	"fmt"
	"strconv"
	"time"
)

var pizzaNum = 0
var pizzaName = ""

func main() {
	//Making three goroutine (lightweight thread) calls and executing parallely
	for threadId := 100; threadId < 103; threadId++ {
		go thread(threadId) //Invoking function in goroutine, this will create asynchronous calls to thread()
	}
	time.Sleep(time.Millisecond * 5000) //Sleeping untill all the goroutine finishes

	//Channels are the pipes which lets goroutine to talk to each other.
	//By default channels are unbuffered. Channels will only accept if there is a corresponding retrieve and also vice-versa.
	channel := make(chan string)                   //Here we are making a channel of type string which is a unbuffered channel
	go func() { channel <- "Passed to Channel" }() //Sending value to channel inside a goroutine
	fmt.Println(<-channel)                         //Reveiving value form channel and print it

	//Buffered channels. To accept more than a single value and store them in buffer to retriev them later, we use buffered channels
	messages := make(chan string, 2) //Here we are making buffered channel which can accept two values.
	//Not more than two values can exits in this. One has to be received before inserting other
	messages <- "buffer1"   //Channel accepting first value
	messages <- "buffer2"   //Channel accepting second value
	fmt.Println(<-messages) //Receving first value from channel
	messages <- "buffer3"   //Channel accepting third value
	fmt.Println(<-messages) //Receving second value from channel
	fmt.Println(<-messages) //Receving third value from channel

	//Channel Synchronization (Using channels for synchronizing goroutines)
	//We can write goroutines in such a way that they are blocked by other goroutine to finish.
	stringChan := make(chan string) //Channel of type string to pass the Pizza Name
	done := make(chan bool)         //Channel of type bool for synchronizing goroutines
	ready := make(chan bool)        //Another Channel of type bool for synchronizing goroutines
	for i := 0; i < 3; i++ {        // Making three Pizzas
		go makeDough(stringChan, done)          //First making a Dough
		go addSauce(stringChan, done)           //Than adding sauce
		go addToppings(stringChan, done, ready) //Than adding toppings
		<-ready                                 //Here we are blocking the execution untill we receive true in ready from addToppings
		pizzaName = <-stringChan                // Receving pizza number from addToppings
		fmt.Println(pizzaName, " Ready")
		time.Sleep(time.Millisecond * 500) //Resting for some time before making a new pizza
	}

	//Channel Directions, We can also specify if a channel can only send or receive value.
	//This increases the type-safety of the program.
	aChan := make(chan string, 1)
	bChan := make(chan string, 1)
	go alert(aChan, bChan, "Alert! ")
	bChan <- "Message!"  //Sending data to bChan to be received in alert
	fmt.Println(<-aChan) //Receiving data from alert

	//Select (It lets you wait for channels to communicate) and Non-Blocking Default
	cChan := make(chan string)
	go func() { //Creating a goroutine which will send "one" on cChan after 1 Sec
		time.Sleep(time.Second * 1)
		cChan <- "one"
	}()
	go func() { //Creating a goroutine which will send "two" on cChan after 2 Sec
		time.Sleep(time.Second * 2)
		cChan <- "two"
	}()
	select { //This select will check if data is received in cChan,i.e, there is a corresponding send, otherwise it will print default.
	case msg := <-cChan: //We don't receive any data on cChan so default will execute.
		fmt.Println("Alert", msg)
	default: //Default is used for implementing non-blocking operations. If there is a send or receive than that case will execute otherwise default will execute without blocking.
		fmt.Println("No Alert")
	}
	for i := 0; i < 2; i++ {
		select { //Here there is no default so it has to wait to receive data in cChan
		case msg := <-cChan:
			fmt.Println("Alert", msg) //After one sec it will set "one" and after two sec "two"
		}
	}

	//Timeout (We can also use timeout to await for some time and send some value). Usefull when talking to external resource or bounding the time of execution
	dChan := make(chan string)
	go func() { //Creating a goroutine which will send "data" in dChan after 2 Sec
		time.Sleep(time.Second * 2)
		dChan <- "data"
	}()
	for i := 0; i < 2; i++ {
		select {
		case <-time.After(time.Second * 1): //First time after waiting for 1 sec it will execute
			fmt.Println("Timeout after Wait")
		case msg := <-dChan:
			fmt.Println("Got ", msg) //In second iteration dChan will get data before timeout
		}
	}

	//Closing a Channel (Once we are done with our task we can close channel, this will also be communicated to the receiver)
	work := make(chan int)
	intimate := make(chan bool)
	go func() {
		for {
			i, more := <-work //Here i will receive value from work. And "more" will be true or false depanding on channel open or closed, or in case of buffered channel if value exist or not.
			if more {
				fmt.Println("Doing Work:", i)
			} else {
				fmt.Println("No Work Received: Vaule of i ", i) //Default value will be 0 of i for int chan if no value or closed chan.
				intimate <- true
				return
			}
		}
	}()
	for j := 1; j <= 3; j++ { //Sending work three times
		work <- j
	}
	close(work) //Closing the Channel
	<-intimate
	fmt.Println("All Work Done")

	//Ranging over Channels
	queue := make(chan string, 3) //Making a buffered channel
	queue <- "one"                //Passing values in the channel
	queue <- "two"
	queue <- "three"
	close(queue)              //Then closing the channel
	for elem := range queue { //This for will iterate till value exist in queue
		fmt.Println(elem)
	}

	//Rate Limiting, this is to control the rate of traffic sent or received.
	//First Example
	requests := make(chan int, 5) //Here we make a buffered chan to handle the incoming requests
	for i := 1; i <= 5; i++ {     //Passed five values to it
		requests <- i
	}
	close(requests)

	//We create a channel which will receive value every 500 millisecs
	limiter := time.Tick(time.Millisecond * 500)

	//Now we retrive values on requests and also we block the execution using limiter
	for req := range requests {
		//Here limiter will get value after every 500 millisecs, so this line will block till then
		//This is how we limit incoming request by 1 per 500 millisecs
		<-limiter
		fmt.Println("Equal Limiting - Request: ", req, "Current Time: ", time.Now())
	}

	//Second Exampple
	//Some time we want to allow the limit to increase and process more requests.
	//We can achive this by using buffered limiter channel
	burstyLimiter := make(chan time.Time, 5)

	//Now initially we send 5 values in channel to burst the limit
	for i := 0; i < 5; i++ {
		burstyLimiter <- time.Now()
	}

	//Aftre that this go routine will keep on sending value over burstyLimiter channel after every 500 millisecs
	go func() {
		for t := range time.Tick(time.Millisecond * 500) {
			burstyLimiter <- t
		}
	}()

	//Now we create a channel from where we are getting requests, we pass 10 request to it
	burstyRequests := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	//Here we retrive values from burstyRequests
	for req := range burstyRequests {
		//Now burstyLimiter already has five buffered values
		//So this code will not block for 5 iterations and it will burst the request processing limit
		//After five iterations it will keep on getting values every 500 millisecs
		<-burstyLimiter
		fmt.Println("Bursty Limiting - Request: ", req, "Current Time: ", time.Now())
	}
}

//This function will be called three times using goroutine, all will run concurrently
func thread(threadId int) {
	for i := 1; i <= 5; i++ {
		fmt.Println(threadId, ":", i) //Each seperate function will print from 1 to 5 taking a pause of 1 sec each.
		time.Sleep(time.Millisecond * 1000)
	}
}

func makeDough(stringChan chan string, done chan bool) {
	pizzaNum++                                     //For every new Pizza assigining a number
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum) //For every new Pizza assigining a name with its number
	fmt.Print("Making Dough for ", pizzaName, " and Send for Sauce")
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Print(".")
	}
	fmt.Println()
	done <- true            //Passing true in done which will be received by addSauce
	stringChan <- pizzaName //Passing Pizza Name to addSauce
}

func addSauce(stringChan chan string, done chan bool) {
	<-done                   //Here we are blocking the execution untill makeDough finishes and passes true
	pizzaName = <-stringChan //Receving Pizza Name from makeDough
	fmt.Print("Adding Sauce for ", pizzaName, " and Send for toppings")
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Print(".")
	}
	fmt.Println()
	done <- true            //Now we pass true to addToppings
	stringChan <- pizzaName //Passing Pizza Name to addToppings
}

func addToppings(stringChan chan string, done chan bool, ready chan bool) {
	<-done                   //Here we are blocking the execution untill addSauce finishes and passes true
	pizzaName = <-stringChan //Receving Pizza Name from addSauce
	fmt.Print("Adding Toppings for ", pizzaName, " and Shipping")
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Print(".")
	}
	fmt.Println()

	ready <- true           //Now our Pizza is ready so we pass true which will be received in main
	stringChan <- pizzaName //Passing Pizza Name to main
}

//This function accepts aChan which can only send data and bChan which can only receive data.
func alert(aChan chan<- string, bChan <-chan string, msg string) {
	msg2 := <-bChan     //Here bChan is receiving data from main
	aChan <- msg + msg2 //Here aChan is sending data to main
}
