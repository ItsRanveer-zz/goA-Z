// go11.go

package main

import (
	"fmt"
	t "time" //using alias t instead of time
)

//Assigning println function to a variable so that we can use it instead
var p = fmt.Println

func main() {

	now := t.Now() //This will get the current time according to the time zone set
	p(now)

	//Creating a struct of "time" by passing the year, month, day, hour, min, sec, nsec and location values to "Date"
	then := t.Date(2014, 05, 01, 21, 43, 37, 917096337, t.UTC)
	p(then)        //After creating the struct we can use it
	p(then.Year()) //Different functions of struct for getting different values
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())

	//Comparing times, then with now. These will return a bool value
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	//Finding the duration between times, then with now
	diff := now.Sub(then) //This will subtract then from now will return difference
	p(diff)               //The output will be in hours, minutes and seconds
	p(diff.Hours())       //We can also get it in various units, like only in hours
	p(diff.Minutes())     //Only in Minutes
	p(diff.Seconds())     //Only in Seconds
	p(diff.Nanoseconds()) //Only in Nanoseconds

	//Can also advance time using a duration or move backwards using a '-' before duration
	p(now)            //Current date and time
	p(now.Add(diff))  //Future date and time
	p(now.Add(-diff)) //Past date and time

	//Unix Epoch. Go's 'time' has Unix epoch functions to find the time since unix epoch(1 January 1970)
	current := t.Now()          //This will give current time
	secs := current.Unix()      //This will give time since Unix Epoch till now
	millis := secs * 1000       //There is no function for Millis, but we can use secs to get that
	nanos := current.UnixNano() //This will give time in nanos

	p("Current Time: ", current)
	p("Time In Sec Since Unix Epoch: ", secs)
	p("Time In Millis Since Unix Epoch: ", millis)
	p("Time In Nanos Since Unix Epoch: ", nanos)

	p("Time in Standard Format :", t.Unix(secs, 0))                 //For getting the corresponding time from secs which is in integers
	p("Time in Standard Format including nanos:", t.Unix(0, nanos)) //For getting the corresponding time from nanos

	//Time Formatting and Parsing using pattern-based layouts.
	running := t.Now()

	//String Formatting of Date and Time
	fmt.Printf("String Formatting: %d-%02d-%02dT%02d:%02d:%02d-00:00 \n",
		running.Year(), running.Month(), running.Day(),
		running.Hour(), running.Minute(), running.Second())

	//Format
	p("Current Time in ANSIC Format: ", running.Format(t.ANSIC))                    //Format will format the time in given format, here ANSIC format
	p("Current Time in RFC3339 Format: ", running.Format(t.RFC3339))                //Format will format the time in RFC3339 format
	p("Current Time in given format: ", running.Format("3:04PM"))                   //We can also supply your custom layouts to format the time accordingly
	p("Current Time in given format: ", running.Format("Mon Jan _2 15:04:05 2006")) //Layout must be exactly like this.
	p("Current Time in given format: ", running.Format("2006-01-02T15:04:05.999999-07:00"))

	//Parse
	parsedTime, _ := t.Parse(t.ANSIC, "Thu May 1 21:10:43 2014") //Parse will parse the given time as the given time format, here ANSIC. Error in second args
	p("Parsed Time From ANSIC Format: ", parsedTime)
	parsedTime2, _ := t.Parse(t.RFC3339, "2014-05-01T21:10:43+00:00") //Parse will parse the given time as the given time format, here RFC3339
	p("Parsed Time From RFC3339 Format: ", parsedTime2)
	parsedTime3, _ := t.Parse("3 04 PM", "4 52 PM") //Here Parse will parse the given time as the given pattern "3 04 PM"
	p("Parsed Time From '3 04 PM' Format: ", parsedTime3)
}
