package main

import (
	"docs/playground/files"
	"time"

	"aidanwoods.dev/go-paseto"
)

var key paseto.V4SymmetricKey

func init() {
	key = paseto.NewV4SymmetricKey()
}

func main() {
	now := time.Now()
	token := paseto.NewToken()
	token.SetIssuer("itlg-security")
	token.SetSubject("authenticated-user-1")
	token.SetAudience("itlg-fe")
	token.SetNotBefore(now)
	token.SetIssuedAt(now)
	token.SetExpiration(now.Add(2 * time.Minute))
	token.SetJti("unique-id")

	encrypted := token.V4Encrypt(key, nil)

	files.WriteToFile("paseto.text", encrypted)
	files.WriteToFile("hex.text", key.ExportHex())
}
