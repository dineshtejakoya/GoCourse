package intermediate

import (
	"encoding/base64"
	"fmt"
)

// Text encoding : ascii, utf-8, utf-16
// Data Encoding : Base64, URL Encoding
// File Encoding : Binary Encoding, Text Encoding
func main() {
	//commonly used for transferring binary data into a textual representation using a set of 64 ascii characters
	data := []byte("He~lo, Base64 Encoding")

	//Encode Base 64
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	//Decode from base64
	decode, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error in decoding:", err)
		return
	}
	//it returns a byte slice like [72 101 108 108 111 44 32 66 97 115 101 54 52 32 69 110 99 111 100 105 110 103]
	fmt.Println("Decoded string:", decode)

	fmt.Println("Decoded string:", string(decode))

	//For url safe encoding avoid '/' and '+'
	//suppose in the string if add ~, + will come in encoded value
	//eg: data := []byte("He~lo, Base64 Encoding")  --> encoded value: SGV+bG8sIEJhc2U2NCBFbmNvZGluZw==

	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	//it removes any / or + symbol within the url
	fmt.Println("URL Safe Encoded:", urlSafeEncoded)

}

//base64 encoding is useful in embedding small images or files directly into html/css using data urls
//we can also store binary data in text based formats such as json or xml
//base64 is not encryption
//proper handling of padding which is equal to symbol
