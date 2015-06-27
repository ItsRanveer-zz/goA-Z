// go12.go

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	var match bool
	//Regexpression - Matching the pattern with string
	match, _ = regexp.MatchString("t([a-z]+)t", "test") //Will return bool value if pattern matches a string
	fmt.Println(match)
	//We can also parse and create the Regexp object which have lot of methods implemented
	regExpObj, err := regexp.Compile("t([a-z]+)t") //Compile will parse regex and return Regexp(struct) object is success otherwise nil
	fmt.Println(regExpObj)
	fmt.Println(err) //If no err than nil

	fmt.Println(regExpObj.MatchString("test")) //Calling MatchString of type Regexp to match the string with pattern

	fmt.Println(regExpObj.FindString("pattern teesst test"))      //Returns the first matched string which matches the pattern
	fmt.Println(regExpObj.FindStringIndex("pattern teesst test")) //Returns the start and end of the first matched pattern

	fmt.Println(regExpObj.FindStringSubmatch("pattern teesst test"))      //Returns the first matched string and the submatches withing that match in an array
	fmt.Println(regExpObj.FindStringSubmatchIndex("pattern teesst test")) //Returns the indexs of match and submatch

	fmt.Println(regExpObj.FindAllString("pattern teesst test turtle", -1))      //Returns all matches in an array, when second argument is negative
	fmt.Println(regExpObj.FindAllString("pattern teesst test turtle", 2))       //Here its 2 so it will only match first two and return that
	fmt.Println(regExpObj.FindAllStringIndex("pattern teesst test turtle", -1)) //Returns the index of all string matches, we can change it by changing the second argument

	fmt.Println(regExpObj.FindAllStringSubmatch("pattern teesst test turtle", -1))      //Returns matches and corresponding submatches in an array and than respectively all in another array
	fmt.Println(regExpObj.FindAllStringSubmatchIndex("pattern teesst test turtle", -1)) //Returns index of matches and submatches

	fmt.Println(regExpObj.Match([]byte("test"))) //If we have array of bytes

	regExpObj = regexp.MustCompile("t([a-z]+)t") //MustCompile will panic if not parsed successfully
	fmt.Println(regExpObj)

	fmt.Println(regExpObj.ReplaceAllString("pattern teesst test", "passed")) //Will replace all the match with the given string

	input := []byte("pattern teesst test")                              //Stored a array of bytes
	fmt.Println(string(regExpObj.ReplaceAllFunc(input, bytes.ToUpper))) //Passing a function in second argument will apply on all the matched patterns
}
