// go18.go

package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	//Base64 is used in MIME for encoding and decoding
	data := "abc123!?$*&()'-=@~" //sample string to encode decode
	fmt.Println("Initial data:", data)

	//Go supports both standard and URL-compatible base64.
	//Standard base64 format Encoding
	//EncodeToString accepts bytes, so we cast our data in bytes, and returns encoded string
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded data using Standard base64 format:", sEnc)

	//Standard base64 format Decoding
	//DecodeString accepts string and returns decoded data in bytes and err if any
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println("Decoded data using Standard base64 format:", string(sDec)) //Decoded data in bytes so we cast it to string

	//URL-compatible base64 format Encoding
	//Encoded data will be slight different than Standard format encoded data
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded data using URL-compatible base64 format:", uEnc)

	//URL-compatible base64 format Decoding
	//Decoded data will be same as we got after decoding using Standard format
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println("Decoded data using URL-compatible base64 format:", string(uDec))
}
