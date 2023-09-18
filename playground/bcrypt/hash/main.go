package main

import (
	"docs/playground/files"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash, err := bcrypt.GenerateFromPassword([]byte("P@ssW0rd!*"), 4)
	if err != nil {
		log.Fatal(err)
	}
	err = files.WriteToFile("hash.text", string(hash))
	if err != nil {
		log.Fatal(err)
	}
}
