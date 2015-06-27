// go13.go

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func main() {
	//Encoding of some datatypes to JSON strings
	bolB, _ := json.Marshal(true)     //Marshal returns encoding(type []uint8) and error
	fmt.Println(reflect.TypeOf(bolB)) //bolB will be of type []uint8
	fmt.Println(bolB)                 //The character encode value of boolean true
	fmt.Println(string(bolB))         //When can use string function to get the actual value from JSON string

	intB, _ := json.Marshal(1) //Encoding int 1
	fmt.Println(intB)          //The character encode value of int 1
	fmt.Println(string(intB))  //Actual value

	fltB, _ := json.Marshal(1.23) //Encoding float 1.23
	fmt.Println(fltB)             //The character encode value of float 1.23
	fmt.Println(string(fltB))     //Value

	strB, _ := json.Marshal("hello") //Encoding string "hello"
	fmt.Println(strB)                //The character encode value of string "hello".When the value is encoded it will encode even the "" also for a string
	fmt.Println(string(strB))        //Value

	//Encoding of slice to JSON arrays
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)     //Encoding slice
	fmt.Println(reflect.TypeOf(slcB)) //Type of encoded value
	fmt.Println(slcB)                 //Character encode of slice. When slice is encoded it will also encode special characters
	fmt.Println(string(slcB))         //Value

	//Encoding of map
	mapD := map[string]int{"red": 3, "blue": 4}
	mapB, _ := json.Marshal(mapD) //Encoding map
	fmt.Println(mapB)             //Character encode of map
	fmt.Println(string(mapB))     //Value

	//Encoding custom types struct
	res1D := Response1{Page: 1, Colors: []string{"red", "blue", "green"}}
	res1B, _ := json.Marshal(res1D) //Encoding struct instance
	fmt.Println(res1B)              //Character encode of struct
	fmt.Println(string(res1B))      //Value

	res2D := Response2{Page: 1, Colors: []string{"red", "blue", "green"}}
	res2B, _ := json.Marshal(res2D) //Encoding struct instance
	fmt.Println(res2B)              //Character encode of struct
	fmt.Println(string(res2B))      //Value with custome JSON tags which was given while declaration

	//Streaming JSON encodings directly to Stdout
	enc := json.NewEncoder(os.Stdout) //Creating a encoder that will write to Stdout
	d := "hello"
	enc.Encode(d) //This will encode d and write to Stdout

	//Decoding
	var data map[string]interface{}                               //Defining a map of string to arbitrary data types to put decoded data
	byt := []byte(`{"num":6.13,"strs":["a","b"],"word":"hello"}`) //Creating a byte to decode
	if err := json.Unmarshal(byt, &data); err != nil {            //Unmarshal will decode byt and put it in dat. Will also return error if occurred
		panic(err)
	}
	fmt.Println(data)         //The map
	fmt.Println(data["num"])  //num in map
	fmt.Println(data["strs"]) //strs in map
	fmt.Println(data["word"]) //word in map

	strs := data["strs"].([]interface{}) //For accessing nested data we have to do casting
	fmt.Println(strs[0], strs[1])

	byt2 := []byte(`{"pg": 1, "clrs": ["red", "blue"]}`) //Creating a byte to decode
	res := Response2{}                                   //Creating a instance to decode to
	json.Unmarshal(byt2, &res)                           //Decoding byt2 to res type
	fmt.Println(res)                                     //Value in res type
	fmt.Println(res.Page)                                //Value in its variabels
	fmt.Println(res.Colors)
}

//Defining structs for encoding and decoding of custom types
type Response1 struct {
	Page   int
	Colors []string
}

type Response2 struct {
	Page   int      `json:"pg"` //Creating JSON tags to customize encoded JSON keys
	Colors []string `json:"clrs"`
}
