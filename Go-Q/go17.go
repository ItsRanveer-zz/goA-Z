// go17.go

package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {

	//Hashing, for security purpose we convert our data into hashes
	//Go has several hash function for creating hashes inside crypto/* package.
	secretMsg := "Secret Text" //Normal text which needs to be hashed

	//Hash by SHA1 which is one of the function for creating hashes.
	h1 := sha1.New() //First we create a new hash.Hash

	h1.Write([]byte(secretMsg)) //Passing msg to Write in bytes

	finalSHA1Hash := h1.Sum(nil) //Final SHA1 Hash will be stored in byte slice

	fmt.Println("Actual Message:", secretMsg)
	fmt.Println("SHA1 Hash in Bytes: ", finalSHA1Hash)
	fmt.Printf("SHA1 Hash in Hex: %x\n", finalSHA1Hash)

	//Hash by MD5 which is one of the function for creating hashes.
	h2 := md5.New() //First we create a new hash.Hash

	h2.Write([]byte(secretMsg)) //Passing msg to Write in bytes

	finalMD5Hash := h2.Sum(nil) //Final MD5 Hash will be stored in byte slice

	fmt.Println("MD5 Hash in Bytes: ", finalMD5Hash)
	fmt.Printf("MD5 Hash in Hex: %x\n", finalMD5Hash)
}
