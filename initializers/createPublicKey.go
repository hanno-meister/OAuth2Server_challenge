package initializers

import (
	"github.com/lestrrat-go/jwx/jwk"
)

var PublicKey interface{}
var JwkKey jwk.Key

func CreateJwkKey() {
	// Create public key from private key
	var err error
	PublicKey, err = jwk.PublicRawKeyOf(PrivateKey)

	if err != nil {
		panic("Failed to extract public key")
	}

	// Create jwk.Key from public key
	JwkKey, err = jwk.New(PublicKey)

	if err != nil {
		panic("Failed to create jwk Key!")
	}

}
