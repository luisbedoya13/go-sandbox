package main

import (
	"docs/playground/files"
	"fmt"
	"log"

	"aidanwoods.dev/go-paseto"
)

func main() {
	hex, err := files.ReadFromFile("hex.text")
	if err != nil {
		log.Fatal(err)
	}
	key, err := paseto.V4SymmetricKeyFromHex(hex)
	if err != nil {
		log.Fatal(err)
	}
	parser := paseto.NewParser()
	encrypted, err := files.ReadFromFile("paseto.text")
	if err != nil {
		log.Fatal(err)
	}
	token, err := parser.ParseV4Local(key, encrypted, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(token.Claims())
}
