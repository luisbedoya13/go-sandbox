package main

import (
	"log"

	"github.com/luisbedoya13/sandbox/utils"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash, err := bcrypt.GenerateFromPassword([]byte("P@ssW0rd!*"), 4)
	if err != nil {
		log.Fatal(err)
	}
	err = utils.WriteToFile("hash.text", string(hash))
	if err != nil {
		log.Fatal(err)
	}
}
