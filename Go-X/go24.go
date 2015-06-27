// go24.go
//Maintaining States

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//Thread Pools, to manage state, properly distribute data in go
	//We will run concurrent instances of our go routine(worker) using channel directions
	jobs := make(chan int, 9)    //This buffered channel is to send the job to worker
	results := make(chan int, 9) //This buffered channel is to receive the result from our worker

	//Now we will start three instances of our worker
	//All workers will be blocked untill we send jobs to worker
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results) //worker call
	}

	//Now we will send jobs to our three workers which are waiting for jobs to receive work
	//All three workers will receive work simultaneously through buffered channel
	//So 9 jobs will be sent to buffered channel and each worker will receive 3 jobs each
	for j := 1; j <= 9; j++ {
		jobs <- j
	}

	//After we are done we will close our jobs channel
	close(jobs)

	//Here we will get all the results
	for a := 1; a <= 9; a++ {
		<-results
	}

	//Atomic Counters, another way to manage state in go.
	var counter uint64 = 0 //First we take positive counter

	//Starting 50 goroutines which will all increment counter
	for i := 0; i < 50; i++ {
		go func() {
			//Each goroutine will increment counter after a time span
			for {
				atomic.AddUint64(&counter, 1) //AddUint64 adds 1 to counter.
				runtime.Gosched()             //Gosched yields the processor from current goroutine and allow other goroutine to proceed
			}
		}()
	}

	//Will sleep for a sec so that all the increment on counters finishes
	time.Sleep(time.Second)

	//To get the final value of counter use LoadUint64
	counterFinal := atomic.LoadUint64(&counter) //Pass the address of counter from where value to fetch to LoadUint64
	fmt.Println("Final Counter Value: ", counterFinal)

	//Mutex-Synchronized State, another way to manage complex state
	//In this we do explicit locking with mutexes to synchronize access to shared state across multiple goroutines

	var state = make(map[int]int) //Here we take a map as a state which needs to be shared among goroutines and also the counter ops which is shared

	//Mutex for synchronize access to state.
	var mutex = &sync.Mutex{}

	//Ops will count number of operations performed against the state
	var ops int64 = 0

	//To simulate concurrent reads we start 100 goroutines
	for r := 0; r < 100; r++ {
		go func() {
			total := 0 //For each goroutine we will keep on suming up the values of state
			//Each goroutine will keep on reading data after every time span
			for {
				key := rand.Intn(5)      //Intn will return random integer which we will use as key to access our state
				mutex.Lock()             //Now we lock the mutex to ensure exclusive access to the state
				total += state[key]      //Read the value of state at that key and add to total
				mutex.Unlock()           //Now unlock the mutex after we finished read
				atomic.AddInt64(&ops, 1) //Now add 1 to ops counter which is counting operations performed on state
				runtime.Gosched()        //Gosched yields the processor from current goroutine and allow other goroutine to proceed
			}
		}()
	}

	//To simulate concurrent writes we start 10 goroutines
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)      //Getting Random key to write the value
				val := rand.Intn(100)    //Generating random value to put
				mutex.Lock()             //Locking mutex on the state
				state[key] = val         //Putting value in state
				mutex.Unlock()           //Unlock the mutex
				atomic.AddInt64(&ops, 1) //Now add 1 to ops counter which is counting operations performed on state
				runtime.Gosched()        //Gosched yields the processor from current goroutine and allow other goroutine
			}
		}()
	}

	//Sleep for a sec to let finish all goroutines
	time.Sleep(time.Second)

	//Get the final value of ops counter
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("Final Ops Value: ", opsFinal)

	//Now we lock the mutex and collect the final state of our state map and then unlock it
	mutex.Lock()
	fmt.Println("State Map:", state)
	mutex.Unlock()

	//We can also maintain state by only using goroutines and channels by using built-in synchronization features of goroutines and channels
	//If state will be owned by a single goroutine, this will guarantee that the data is never corrupted with concurrent access
	//To read and write the data, goroutines will send messages to owning goroutines and will receive corresponding replies. Struct readOp and writeOp implements that.

	var operations int64 = 0 //Ops will count number of operations performed against the state

	reads := make(chan *readOp)   //Read channel of type readOp to issue read request by other goroutines
	writes := make(chan *writeOp) //Write channel of type writeOp to issue write request by other goroutines

	//Now we make a goroutine which will own the state which is a map - stateMap, and it is private to the stateful goroutine
	go func() {
		var stateMap = make(map[int]int) //private state map
		for {
			//This goroutine repeatedly selects on the reads and writes channels and than send response on respective resp channel
			select {
			case read := <-reads: //If there is a value on read channel than this will execute
				read.resp <- stateMap[read.key] //Here we send the response on resp channel of readOp

			case write := <-writes: //If there is a value on write channel than this will execute
				stateMap[write.key] = write.val
				write.resp <- true //Here we send the response on resp channel of writeOp
			}
		}

	}()

	//Here we start 100 goroutines which issues reads to the state owning goroutine using reads channel
	for r := 0; r < 100; r++ {
		go func() {
			for {
				//For reading we need to construct a readOp with the key to be read and resp channel to get the response after read
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read                   //Sending value over reads channel to be received by state owning goroutine
				<-read.resp                     //Receiving the response from the state owning goroutine
				atomic.AddInt64(&operations, 1) //Now add 1 to operations counter which is counting operations performed on state
			}
		}()
	}

	//Here we start 10 goroutines which issues writes to the state owning goroutine using writes channel
	for w := 0; w < 10; w++ {
		go func() {
			for {
				//For writing we need to construct a writeOp with the key and value to write and resp channel to get the response after write
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write                 //Sending value over writes channel to be received by state owning goroutine
				<-write.resp                    //Receiving the response from the state owning goroutine
				atomic.AddInt64(&operations, 1) //Now add 1 to operations counter which is counting operations performed on state
			}
		}()
	}

	time.Sleep(time.Second)                          //Sleep for a sec to let finish all goroutines
	operationsFinal := atomic.LoadInt64(&operations) //Get the final value of operations counter
	fmt.Println("Total Operations Value:", operationsFinal)
}

//This is our worker, which we will run multiple instances
//This has one receive channel jobs and one send channel results
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs { //Untill all the data taken from jobs we will keep on iterating
		fmt.Println("Worker", id, " is Processing job ", j)
		time.Sleep(time.Second) //Sleeping for a sec to do the work
		results <- j            //After doing work we will send result to buffered channel results
	}
}

//This struct is to request for a read and than response from the owner goroutine
//For reading we send a key and get the response in resp channel
type readOp struct {
	key  int
	resp chan int
}

//This struct is to request for a write and than response from the owner goroutine
//For writing we send the key and value and get the response in resp channel
type writeOp struct {
	key  int
	val  int
	resp chan bool
}
