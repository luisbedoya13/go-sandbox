package main

import (
	"bufio"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	file, err := os.Open("./playground/examples/hash.text")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if scanner.Scan() {
		bcrypt.CompareHashAndPassword([]byte(scanner.Text()), []byte("P@ssW0rd!"))
		if err == nil {
			log.Println("It works")
		} else {
			log.Fatal(err)
		}
	}
}
