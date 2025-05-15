package utils

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwe"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/spf13/viper"
)

func DecryptJwe(message string) (*string, error) {
	privKey := viper.GetString("secret.inventory.private")
	if privKey == "" {
		return nil, fmt.Errorf("private key not found")
	}

	block, _ := pem.Decode([]byte(privKey))

	ecKey, _ := x509.ParseECPrivateKey(block.Bytes)

	key, err := jwk.Import(ecKey)
	if err != nil {
		return nil, fmt.Errorf("failed to import private key with error: %v", err)
	}

	decrypted, err := jwe.Decrypt([]byte(message), jwe.WithKey(jwa.ECDH_ES(), key))
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt message with error: %v", err)
	}

	text := string(decrypted)

	return &text, nil
}
