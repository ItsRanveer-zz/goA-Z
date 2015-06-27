// go3.go

package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Arrays. These are fixed size data structures in Go. Length of array and type of elements both defines arrays type.
	var favNums [5]float64                                                       //Only defining Array, by default value of the elements of the float and int array will be 0, string array will be empty
	favNums[0] = 163                                                             //Setting value at index 0 of an array
	fmt.Printf("%T - %f - %T - %f \n", favNums, favNums, favNums[0], favNums[0]) //Printing type of array, array itself, type of element at 0 index and value at 0 index

	favNums2 := [5]float64{11, 22, 33, 44, 55} //Defining and initializing at the same time
	fmt.Println(favNums2)

	favNums2 = favNums                          //We can assign same type(same length and element type) array to another array.
	fmt.Println(favNums2, " - ", len(favNums2)) // len function is used to find the length of an array

	//Using range with arrays, range will iterates over whole array and return index and the value at that index
	for index, value := range favNums2 {
		fmt.Println(index, value)
	}

	//Arrays can be used to build multi dimensional data structures
	twoDarray := [2][2]int{{52, 56}, {42, 63}}
	fmt.Println("2D Array - ", twoDarray)

	//Slices (It is a reference type. It is like an array but it will actually hold a reference to an array. You don't define the size and it will dynamically resize)
	var numSlice []int                                                               //Definig slice
	numSlice = []int{22, 59, 52, 65, 68, 55}                                         //Initializing slice
	fmt.Printf("%T - %d - %T - %d \n", numSlice, numSlice, numSlice[2], numSlice[2]) //Printing type of slice, full slice, type at index 2, value at index 2

	numSlice2 := numSlice[3:5] // Creating slice from another slice. [:4] [2:] [:]
	fmt.Println(numSlice2)     // Will output this [65 68]
	numSlice3 := favNums2[:3]  // Creating slice from an array
	fmt.Println(numSlice3)     // Will output this [163 0 0]

	numSlice4 := make([]int, 5, 10) // We can also use make to create slice with following parameters (type, default value of 0 to first n element, max size)
	fmt.Println(numSlice4)          // Default value will be 0 for first 5 elements. Here length will be 5 but capacity will be 10 of this slice
	fmt.Println(len(numSlice4))     // Length of a Slice
	fmt.Println(cap(numSlice4))     // Capacity of a Slice

	copy(numSlice4, numSlice2) // For copying one slice to otherwe use copy (here copying numSlice2 to numSlice4)
	fmt.Println(numSlice4)

	numSlice2 = numSlice4 // We can also assign large slice to a small one, it will dynamically increase the size of samller one
	fmt.Println(numSlice2)
	numSlice2 = append(numSlice2, 5, 6, 55) // Can also use append to add change item in slice. First argument will always be a slice and rest will be values of slice type
	fmt.Println(numSlice2)
	numSlice2 = append(numSlice, numSlice4...) // This is also equivalent to append(numSlice, numSlice4[0], numSlice4[1], numSlice4[2], numSlice4[3], numSlice4[4])
	fmt.Println(numSlice2)

	//Slices can also be used to build multi dimensional data structures
	twoDslice := [][]int{{22}, {42, 63}, {55, 54, 95}}
	fmt.Println("2D Slice - ", twoDslice)

	//Maps (Reference type. It is a collection of Key Value pairs. No order maintained)
	//var mapAge map[string]int .  //It can be defined like this but will be a nill map with length 0 and value 0 for any key.
	//Also if we initialize it will give runtime panic cause there is no referenced map to which it is referencing.
	mapAge := make(map[string]int) //Define like this. It will allocates and initialize a hash map data structure and returns a map value that points to it.
	mapAge["Hari"] = 42
	mapAge["Mahesh"] = 66
	mapAge["Babu"] = 34
	fmt.Println("Hari's Age - ", mapAge["Hari"])                  // Finding value of a key in map
	fmt.Println("Length of Map - ", len(mapAge))                  // Length of map or number of key value pairs in map
	fmt.Println(reflect.TypeOf(mapAge), "-", mapAge)              // Printing type of map and full map
	delete(mapAge, "Mahesh")                                      // For deleting a value from a map. (Map Name, key)
	fmt.Println(mapAge)                                           // Map after deleting a value
	mapAge2 := map[string]int{"Sam": 42, "Scot": 56, "Larry": 41} // Can also be defined and initialized in same line
	fmt.Println(mapAge2)
	mapAge3 := mapAge //Creating one map from another
	fmt.Println(mapAge3)
	//Using range with maps, range will iterates over whole map and return key and the value of that key
	for key, value := range mapAge3 {
		fmt.Printf("%s - %d\n", key, value)
	}
}
