package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash, err := bcrypt.GenerateFromPassword([]byte("P@ssW0rd!"), 4)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile("./playground/examples/hash.text", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Fprint(f, string(hash))
}
