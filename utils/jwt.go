package utils

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"github.com/spf13/viper"
)

func DecryptJwt(message string) (*map[string]string, error) {
	pubString := viper.GetString("secret.sso.public")
	if pubString == "" {
		return nil, fmt.Errorf("public key not found")
	}

	block, _ := pem.Decode([]byte(pubString))

	ecKey, _ := x509.ParsePKIXPublicKey(block.Bytes)

	key, err := jwk.PublicRawKeyOf(ecKey)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse([]byte(message), jwt.WithKey(jwa.ES256(), key), jwt.WithValidate(true))
	if err != nil {
		return nil, err
	}

	keys := token.Keys()
	payloads := make(map[string]string)

	for _, v := range keys {
		var data string

		token.Get(v, &data)

		payloads[v] = data
	}

	return &payloads, nil
}
