// go9

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//File I/O (Creating, Opening, Writing, Closing, Reading)

	//Creating a File
	textfile, err := os.Create("text.txt") //Will create a new file called "text.txt" everytime in current directory, os.Create will return a os.File value
	defer textfile.Close()                 //To close the file at the end of function

	//Opening a File
	//textfile, err := os.OpenFile("text.txt", os.O_APPEND|os.O_WRONLY, 0600) //Can also use this if already have a file instead of creating a new one, returns a os.File value

	check(err) //To log a error if occurs while creating a file

	//Writing to File
	textfile.WriteString("Hello, Wassup\n")                            //WriteString will write the string passed to the file which we opened
	numberByte, err := textfile.WriteString("How is your day going\n") //WriteString will also return number of bytes written to file and err if occurred while writing
	fmt.Printf("%d bytes written to file\n", numberByte)

	dataStream := []byte{116, 101, 120, 116, 10}   //Instead of writing a string to file we will make a byteStream and pass it to Write, here 5 bytes when converted will be "text\n"
	numberBytes, err := textfile.Write(dataStream) //Write will write passed byteStream in the file and return the number of bytes written and err if occured
	check(err)
	fmt.Printf("%d bytes written to file\n", numberBytes)

	textfile.Sync() //Sync will flush the data from main memory to disk so that it becomes stable

	writer := bufio.NewWriter(textfile)                       //buffer input output provides functions for small reads and writes, here it will return a new writer on file
	numberBytes0, err := writer.WriteString("Buffered Write") //WriteString on writer will write the string passed in buffered way and return no. of bytes written and err if not able to write of written is less than what was send
	check(err)
	fmt.Printf("%d bytes written to file\n", numberBytes0)

	writer.Flush() //Flush ensure all buffered data written to the Writer

	//Reading a File
	stream, err := ioutil.ReadFile("text.txt") //ReadFile to read all the content of the file named text.txt and return the content as stream
	check(err)
	fileContent := string(stream) //Readfile returns []uint8 stream. We have to convert it to string
	fmt.Println(fileContent)

	//Some other ways
	//Opening a File
	file, err := os.Open("text.txt") //This also opens the file but has only READONLY flag
	check(err)
	defer file.Close() //Closing the file at the end of the program

	//Read a file first five byte
	byteStream := make([]byte, 5)              //Make a stream in which we have to put the readed content
	numberBytes1, err := file.Read(byteStream) //Read will read the file and put the stream in passed byteStream(upto len of byteStream, so rest will be ignored) and return number of bytes read
	check(err)
	fmt.Printf("%d bytes: %s\n", numberBytes1, string(byteStream))

	//Read a file content at some location
	offset, err := file.Seek(8, 0) //Seek will set new offset for next Read and Write on that file, here its 8 so next read or write will happen at 8th byte
	check(err)
	byteStream2 := make([]byte, 2)
	numberBytes2, err := file.Read(byteStream2) //Reading two bytes a new offset which is 8
	check(err)
	fmt.Printf("%d bytes from offset @%d: %s\n", numberBytes2, offset, string(byteStream2))

	numberBytes3, err := file.Read(byteStream2) //After the above read of two bytes again new offset is set to 8+2 10
	check(err)
	fmt.Printf("%d bytes from offset @%d: %s\n", numberBytes3, offset, string(byteStream2))

	offset2, err := file.Seek(0, 0) //Once again we will set the offset to start
	check(err)
	byteStream3 := make([]byte, 5)
	numberBytes4, err := io.ReadAtLeast(file, byteStream3, 2) //ReadAtLeast is more robust way of reading a file, here we are reading a min of 2 bytes, in less bytes than error will occur
	check(err)
	fmt.Printf("%d bytes from offset @%d: %s\n", numberBytes4, offset2, string(byteStream3))

	file.Seek(0, 0)                    //Again setting offset to 0
	reader := bufio.NewReader(file)    //buffer input output provides some functions for small reads and writes, here it will return a new reader on file
	byteStream4, err := reader.Peek(5) //Peek will return a byteStream till the passed value, if nothing to read or less to read than it will also retun error
	check(err)
	fmt.Printf("5 bytes: %s\n", string(byteStream4))

	//Writing to File
	dataStream2 := []byte("Konichiwa\nWorld!")             //Creating byte stream to write
	err = ioutil.WriteFile("text2.txt", dataStream2, 0644) //Write file will create a file if not exits with 0644 permission, write the datastream into it and close it
	check(err)
}

//To log a error if occurs while creating a file
func check(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
		//panic(e) //We can also Panic
	}
}
