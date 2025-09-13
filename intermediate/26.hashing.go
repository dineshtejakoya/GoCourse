package intermediate

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

//hashing is irreversible
//fast to compute

//real life use cases: when we are entering our password the pwd is not entered directly into the database
//the pwd is usually converted to a hash value and the hash value is stored in a database

// Our APIs are designed to not directly match the input from the pwd field from the value stored in the database,
// it first converts using a hash function, and then matches it with the value that is stored in the database
// it is to maintain the integrity of the data
// is used in quick data retrieval
// hash functions are irreversible
// 1.Good hash functions are designed to minimise collissions(different inputs produce the same hash output)
// 2.avalance effect
func main() {
	password := "password123"
	// //sum256 function returns checksum of the data
	// //this will only accept byte slice, we need to convert this password string to byteslice
	// hash := sha256.Sum256([]byte(password))
	// hash512 := sha512.Sum512([]byte(password))
	// fmt.Println(password)

	// //this byteslice displays the ascii values of the characters in string, but not the characters that we enter in our password
	// fmt.Println(hash)
	// fmt.Println(hash512)
	// fmt.Printf("SHA-256 Hash hex val: %x\n", hash)
	// fmt.Printf("SHA-512 Hash hex val: %x\n", hash512)

	//Salting adds an extra layer of security by combining the pwd with a unique random value
	//The practice of salting helps protect against dictionary attacks, rainbowtable attacks
	//so, salt is a value added to the password before hashing,its purpose it to ensure even if two users has the same password their hashvalues will be different due to different salts
	salt, err := generateSalt()
	fmt.Println("Original Salt:", salt)
	fmt.Printf("Original Salt: %x\n", salt)
	if err != nil {
		fmt.Println("Error Generating salt:", err)
		return
	}

	//Hash password with salt
	signUpHash := hashPassword(password, salt)

	//Store the salt and password in a database, just printing as of now
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt:", saltStr)     //simulate as storing in database
	fmt.Println("Hash: ", signUpHash) //simulate as storing in database

	hashOriginalPassword := sha256.Sum256([]byte(password))
	fmt.Println("Hash of just the password string without salt:", base64.StdEncoding.EncodeToString(hashOriginalPassword[:]))

	//verify
	//Decode the saltStr
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Unable to decode salt", err)
		return
	}
	loginHash := hashPassword(password, decodedSalt)

	//compare the signupHash with Loginhash
	if signUpHash == loginHash {
		fmt.Println("Password is correct. You are logged in")
	} else {
		fmt.Println("Login Failed. Please check user credentials")
	}
}

// Generate Salt
func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	//now, we use the io package readfull function, it'll read exactly the length of the byteslice from the given reader into the salt slice
	//here it returns the no. of bytes it read and an error
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// Function to hash password
func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
