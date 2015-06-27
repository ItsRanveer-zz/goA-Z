// go19.go

package main

import (
	"crypto/sha1"
	"fmt"
	"net"
	"net/url"
)

func main() {
	p := fmt.Println

	//Here is a URL which contains scheme, authentication info, host, port, path, query params, and query fragment.
	rawUrl := "mongodb://user:pass@openshift.com:27017/home.html?sP=Name&sV=Veer&sO=eq#3"

	//Parse from url package will parse raw url and returns the URL structure or err if occured
	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		panic(err)
	}

	p(parsedURL) //Parsed URL

	//Accessing Properties in the URL.
	p("Scheme:", parsedURL.Scheme)                         //Scheme will give Scheme in your URL
	p("User:", parsedURL.User)                             //User contains userinfo - Username and Password in your URL
	p("Username:", parsedURL.User.Username())              //To get just Username we can use Username on User
	password, _ := parsedURL.User.Password()               //Password will return password and bool value specifying that a pass is set not
	p("Password:", password)                               //Password
	p("Host:", parsedURL.Host)                             //Host contains both hostname and port
	hostName, port, _ := net.SplitHostPort(parsedURL.Host) //SplitHostPort will seperate the hostname and port
	p("HostName:", hostName)                               //Host
	p("Port:", port)                                       //Port
	p("Path:", parsedURL.Path)                             //Path
	p("RawQuery:", parsedURL.RawQuery)                     //Whole Query Part in url can get by RawQuery
	parsedQuery, _ := url.ParseQuery(parsedURL.RawQuery)   //To fetch the parameters in query we have to parse using ParseQuery which returns a map
	p("ParsedQuery Map:", parsedQuery)                     //Parsed Query in a map
	p("Query Key sP:", parsedQuery["sP"][0])               //Fetching the key values from the parsed query
	p("Query Key sV:", parsedQuery["sV"][0])
	p("Query Key sO:", parsedQuery["sO"][0])
	p("Fragment:", parsedURL.Fragment) //Any fragment on the page
}
