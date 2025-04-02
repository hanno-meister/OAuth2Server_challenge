package initializers

import (
	"crypto/rand"
	"crypto/rsa"
)

var PrivateKey *rsa.PrivateKey

// Generate random RSA private key
func CreatePrivateKey() {
	var err error
	PrivateKey, err = rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		panic("Failed to generate RSA private key!")
	}
}
