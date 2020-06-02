package auth

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/pkg/errors"
)

// Algorithm to be used for private key
const algorithm = "RS256"

// KeyGen creates an x509 private key for signing auth tokens.
func KeyGen() ([]byte, error) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return []byte{}, errors.Wrap(err, "generating keys")
	}

	block := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	buf := new(bytes.Buffer)
	if err := pem.Encode(buf, &block); err != nil {
		return []byte{}, errors.Wrap(err, "encoding to private file")
	}

	return buf.Bytes(), nil
}
