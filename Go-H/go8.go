// go8.go

package main

import (
	"fmt"
	"sort"
)

func main() {
	//Sorting
	strs := []string{"abc", "!.", "abb", "aa.c", "12"} //Slice of strings to be sorted
	sort.Strings(strs)                                 //Strings used to sort a slice of strings lexicographically
	fmt.Println("Strings:", strs)                      //It won't return a new slice rather changes the given slice
	ints := []int{13, 3, 11}                           //Slice of ints to be sorted
	sort.Ints(ints)                                    //Ints used to sort ints
	fmt.Println("Ints:   ", ints)

	//Custom Sort by Length of string
	fruits := []string{"aaaaa", "cccccc", "bbb"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}

//Created our own custom []string type to use with our custom function for sorting
type ByLength []string

//This Function will return length of string
func (s ByLength) Len() int {
	return len(s)
}

//Function use for swapping
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//Function will compare the length of two strings and return true and false
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
