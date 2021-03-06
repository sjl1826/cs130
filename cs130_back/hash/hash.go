package hash

import (
	"log"
	"strings"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// Salt generates a hashed password
func Salt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

// ComparePasswords hashes the plain password and compares it with the hashed password
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	if strings.HasSuffix(os.Args[0], ".test") {
		return (hashedPwd == string(plainPwd)) // don't unhash password if test
	}
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
