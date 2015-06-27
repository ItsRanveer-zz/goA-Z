//go1.go
/*
Go is opensource, statically compiled, statically typed, concurrent, garbage collected,
easy to understand and learn, and fast programming language.

To build and run the program you can open this file in any Go editor (LiteIDE, Eclipse with Goclipse, IntelliJ IDE with Go) or
Either you can use command line.

$ go run go1.go

For building the programs into binaries we use

$ go build go1.go

Now we can run that binary directly

$ ./go1

Also after all these above commented lines you would have got to know how to give single line and multiline comment in Go
*/

/*Package Name deceleration ("go build" command will search for "main" package in the Project, Also there will be only one package,i.e. "main" in any Go Project,
if defined more than one than it will give error "can't load packages")*/
package main

//Importing Packages. These are some of the important packages which comes along go. We are going to use all of them, if we don't use
//then program will not compile. So remeber in Go whatever packages you import or variables you define you have to make use of it
//We can import multiple with this syntax
import (
	"fmt"
	"reflect"
)

//We can also import by single line import syntax
import "strings"

//We Can also define alias to the package name so we can use instead of the package name
//Now we have to use s instead of strconv
import s "strconv"

//Defining main function (There can and should be only one main function otherwise error - "main redeclared" or "undefined main")
func main() {

	//Normal Printing, Here "Hello World!" is a value of type string which we are directly passing to Prinln function and not storing in a variable.
	fmt.Println("Hello World!")   // Println prints to os.Stdout and changes the line after printing, it also returns the length of string and error
	fmt.Print("Hello ")           // Print will not change the line after printing
	fmt.Print("World \t Twice\n") // We can use \n to terminate the line or \t to give space

	//Defining Variables and Constant
	//We can define many value types like string, integers, floats, booleans etc
	var age int              //Defining variable 'age' of type int. If no value given than by-default it will be 0 for integer
	fmt.Println("Age:", age) //Default value 0
	age = 40                 //assigning value 40 to age

	var favNum float32 = 1.6837 //float type. Defining and assigning value. By-default it will be 0.000000 for integer

	//Defining variable with shorthand syntax ":=". Can also use "var randNum = 1". Go will infer the type of the variable.
	randNum := 1

	var myName string = `Ranveer` + " Singh" //Can use "" or `` for strings, default value of string will be empty

	var isOver40 bool = true //default value will be false for a bool type

	//defining multiple variables at a time. Can also use "var num1, num2 = 5, 6"
	var (
		num1 = 1.000
		num2 = .9999
	)

	//defining constant
	const pii float64 = 3.14159265

	//String Formatting. We use Printf for string formatting
	fmt.Printf("Integer: %d \n", age)                           //for integer(base 10)
	fmt.Printf("String: %s \n", myName)                         //for string
	fmt.Printf("String quotes: %q\n", "\"string\"")             //for string with quotes
	fmt.Printf("Float: %f \n", favNum)                          //for float, by default till 6 decimal point
	fmt.Printf("Float 3: %.3f \n", pii)                         //for floating point till 3 position after decimal
	fmt.Printf("Width and Left Allign:|%6d|%6d|\n", 2, 1234)    //for width based formatting, by default the numbers will be right allingned in that certain width
	fmt.Printf("Width and Right Allign:|%-6d|%-6d|\n", 2, 1234) //for width based formatting left alligned, use s for strings in place of d
	fmt.Printf("Binary: %b \n", 100)                            //for binary code
	fmt.Printf("Hex Code: %x \n", 17)                           //for hex code of integer
	fmt.Printf("Hex Code String: %x \n", "hi")                  //for hex code of string with two character o/p per byte of input
	fmt.Printf("Character Code: %c \n", 44)                     //for character code of integer
	fmt.Printf("Boolean: %t \n", isOver40)                      //for boolean
	fmt.Printf("Scientific Notation: %e \n", pii)               //for Scientific Notation with e
	fmt.Printf("Scientific Notation: %E \n", pii)               //for Scientific Notation with E
	fmt.Printf("%T \n", pii)                                    //for finding type of variable or const

	//for finding type of a variable or const we can also use this
	fmt.Println(randNum, "-", reflect.TypeOf(randNum))

	fmt.Println("Floating Point Substraction of two float64 - ", num1-num2)

	//Only String Formatting without printing to os.Stdout.
	str := fmt.Sprintf("Normal %s", "string") //This will format and keep the value in str
	fmt.Println(str)

	//String Functions
	sampleStr := "My name is Ranveer Singh"
	fmt.Println("Contains:", strings.Contains(sampleStr, "in"))      //Returns true or false if contains
	fmt.Println("Index:", strings.Index(sampleStr, "a"))             //Returns the first matached index
	fmt.Println("Count:", strings.Count(sampleStr, "i"))             //Returns count matched in String
	fmt.Println("Replace:", strings.Replace(sampleStr, "n", "x", 2)) //Replaces first 2 "n" with "x" and returns the new string, Put -1 for replace all
	fmt.Println("Repeat:", strings.Repeat("ab", 3))                  //Repeats a string certain times and creates a single string
	fmt.Println("Split:", strings.Split(sampleStr, "n"))             //Splits with "n" and puts in array and returns that array
	fmt.Println("HasPrefix:", strings.HasPrefix(sampleStr, "My"))    //Return bool if string has prefix
	fmt.Println("HasSuffix:", strings.HasSuffix(sampleStr, "gh"))    //Return bool if string has suffix
	fmt.Println("ToLower:", strings.ToLower(sampleStr))              //Changes string to lower case
	fmt.Println("ToUpper:", strings.ToUpper(sampleStr))              //Changes string to upper case
	listOfLetters := []string{"c", "h", "a"}                         //Slice of string
	fmt.Println("Join:", strings.Join(listOfLetters, ", "))          //Joins the slice by "," and returns a single string
	fmt.Println("Length:", len(sampleStr))                           //Length of string
	fmt.Println("CharAt:", sampleStr[4])                             //Returns unicode of character at 4th index of string

	//Type Casting
	randInt := 44
	randFloat := 10.5
	randString := "100"
	randString2 := "250.5"
	fmt.Println(float64(randInt), "-", reflect.TypeOf(float64(randInt))) //Integer to Float64
	fmt.Println(int(randFloat), "-", reflect.TypeOf(int(randFloat)))     //Float64 to Integer
	newInt, _ := s.ParseInt(randString, 0, 64)                           //String to Integer64 (Will give 0 if found any character in string)(Returns two value, other one is error)
	fmt.Println(newInt, "-", reflect.TypeOf(newInt))
	newFloat, _ := s.ParseFloat(randString2, 64) //String to Float64 (Will give 0 if found any character in string)
	fmt.Println(newFloat, "-", reflect.TypeOf(newFloat))
	fmt.Println(string(randInt), "-", reflect.TypeOf(string(randInt)))                                           //string() will consider randInt as character code and convert it to that character ","
	fmt.Println(s.Itoa(randInt), "-", reflect.TypeOf(s.Itoa(randInt)))                                           //Integer to String
	fmt.Println(s.FormatFloat(randFloat, 'f', 6, 64), "-", reflect.TypeOf(s.FormatFloat(randFloat, 'f', 6, 64))) //Float64 to String
}
