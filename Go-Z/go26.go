//go26.go

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	//Web Server
	//For handeling a url pattern we use Handel and HandleFunc
	//For Serving the file from a folder, static in this case we use FileServer
	//Here we define that for default / url serve the static folder, where by default index.html will be loaded for this url.
	http.Handle("/", http.FileServer(http.Dir("./static")))
	//For second page we are calling a handler function
	http.HandleFunc("/secondpage", secondPageHandler)

	log.Println("Listening at 8000...") //This will be printed on the console.
	log.Println("Go ahead and browse localhost:8000")

	//ListenAndServe will start listening on TCP network address at given port, second argument can be a handler to handel request on incoming connections
	err := http.ListenAndServe(":8000", nil)
	//ListenAndServe returns error if occurred
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//Defining Handler Function for Second Page
//ResponseWriter interface is used by an HTTP handler to construct an HTTP response
//Request represents an HTTP request received by a server or to be sent by a client.
func secondPageHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Everyone!\n")         //WriteString will write the string, second args, on the ResponseWriter w
	fmt.Fprintf(w, "Welcome to the second page\n") //Can also use Fprintf for formating string, this also writes the string on the ResponseWriter w
	sometext := "Enjoy Coding in Go"
	fmt.Fprintf(w, sometext)
}
