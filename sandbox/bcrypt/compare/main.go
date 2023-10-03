package main

import (
	"log"

	"github.com/luisbedoya13/sandbox/utils"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash, err := utils.ReadFromFile("hash.text")
	if err != nil {
		log.Fatal(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte("P@ssW0rd!*"))
	if err == nil {
		log.Println("It works")
	} else {
		log.Fatal(err)
	}
}
