package main

import (
	"fmt"
	"log"

	"aidanwoods.dev/go-paseto"
	"github.com/luisbedoya13/sandbox/utils"
)

func main() {
	hex, err := utils.ReadFromFile("hex.text")
	if err != nil {
		log.Fatal(err)
	}
	key, err := paseto.V4SymmetricKeyFromHex(hex)
	if err != nil {
		log.Fatal(err)
	}
	parser := paseto.NewParser()
	encrypted, err := utils.ReadFromFile("paseto.text")
	if err != nil {
		log.Fatal(err)
	}
	token, err := parser.ParseV4Local(key, encrypted, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(token.Claims())
}
