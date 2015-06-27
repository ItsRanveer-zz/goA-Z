//go23.go

package main

import (
	"fmt"
	"time"
)

func main() {
	//Timers (Used to trigger some event after certain time once in the future)
	timer1 := time.NewTimer(time.Second * 1) //Creating a timer which will send value to a channel after 1 sec

	<-timer1.C //Here it will wait on timer1's channel C to receive value and get expired
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second * 1) //Creating one more channel
	go func() {
		<-timer2.C //This will also wait for timer2 to get expired
		fmt.Println("Timer 2 expired")
	}()
	//We could have used Sleep for waiting but in case of timers we can stop it from waiting.
	stop2 := timer2.Stop() //Here we stopped timer2 from waiting, it will return bool true
	if stop2 {
		fmt.Println("Timer 2 stopped") //While timer2 was waiting we stopped it.
	}

	//Tickers (Used to trigger some event repeatedly at regular intervals)
	ticker := time.NewTicker(time.Second * 1) //Here we create a ticker which will send value to channel repeatedly after every 1 sec
	go func() {
		for t := range ticker.C { //This for loop will keep on receiving values from the ticker channel repeatedly after every 1 sec
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(time.Second * 3) //Sleeping for three second and giving some time to ticker for ticking
	ticker.Stop()               //Stoping the ticker, if we won't stop than our ticker will keep ticking.
	fmt.Println("Ticker stopped")
}
