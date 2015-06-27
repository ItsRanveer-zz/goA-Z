//Pop Quiz and Tricks

package main

import (
	"fmt"
)

func main() {
	//First Example
	//A is an array, when giving values we can mention the index where values needs to be stored
	//Here 5 will go at 0th index but second value 1 will go at 4th index, and than 0 will go to next index to that,i.e,5th index
	a := [...]int{5, 4: 1, 0, 2: 3, 2, 1: 4}
	fmt.Println(a) //Output will be [5 4 3 2 1 0]

	//Second Example
	//Making a slice of Foo structs  having three Foo
	orig := []Foo{
		{"1"},
		{"2"},
		{"3"},
	}

	//Create slice of Foo pointers ([]*Foo) from []Foo, pointer of Foo same length as Foo
	ptrs := make([]*Foo, len(orig))

	/*Wrong implementation
	Here we are storing address of a in ptrs slice, so when we retrive value it will give the final value at that address
	for i, a := range orig {
		fmt.Printf("%v %v %p\n", i, a ,&a)
		ptrs[i] = &a
	}*/

	//Correct Implementaion, here we are storing the address of actually foo structs in the ptrs slice
	for i := range orig {
		ptrs[i] = &orig[i]
	}
	fmt.Printf("Values: %v %v %v\n", orig[0], orig[1], orig[2])
	fmt.Printf("Pointers: %v %v %v\n", ptrs[0], ptrs[1], ptrs[2])

	//Third Example, Finding min using one liner
	f(9000, 314)
	f(490, 500)

	//Forth Example
	//Here we make a map and than default value of any new key will be 0, so we increment it and print it
	m := make(map[string]int)
	m["foo"]++
	fmt.Println("Value at key foo", m["foo"])

	//Fifth Example
	i, j, k := 1, 2, 3

	switch i > j {
	case i > j:
		fmt.Println("i > j")
	case i <= j:
		fmt.Println("i <= j")
	case i <= k:
		fmt.Println("i <= k")
	}
}

//Defining a struct Foo which will have one property
type Foo struct {
	Bar string
}

func f(a, b int) {
	//In the last argument we take a map, put the values and than lookup for evaluated bool value (true or false) from [a<b] in the map
	fmt.Printf("Using Map: The min of %d and %d is %d\n", a, b, map[bool]int{true: a, false: b}[a < b])
	//Another method
	fmt.Printf("Using Copy: The min of %d and %d is %d\n", a, b, copy(make([]struct{}, b), make([]struct{}, a)))
}
