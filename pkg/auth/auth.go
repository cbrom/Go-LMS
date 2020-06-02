package auth

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

var secretKey []byte

// KeyFunc generates a new private public key pairs
type KeyFunc func(keyID string) (*rsa.PublicKey, error)

// NewKeyFunc is a multiple implementation of KeyFunc that supports a map of keys
func NewKeyFunc(keys map[string]*PrivateKey) KeyFunc {
	return func(kid string) (*rsa.PublicKey, error) {
		key, ok := keys[kid]
		if !ok {
			return nil, fmt.Errorf("unrecognized key id %q", kid)
		}
		return key.Public().(*rsa.PublicKey), nil
	}
}

func init() {
	pwd, _ := os.Getwd()
	keyPath := pwd + "/api/jwtsecret.key"

	key, readErr := ioutil.ReadFile(keyPath)
	if readErr != nil {
		fmt.Println(keyPath)
		// panic("failed to load secret key file")
	}

	secretKey = key
}

// Authenticator defines an authenticator struct
type Authenticator struct {
	privateKey *PrivateKey
	keyID      string
	algorithm  string
	kf         KeyFunc
	parser     *jwt.Parser
}

// PrivateKey defines private key struct
type PrivateKey struct {
	*rsa.PrivateKey
	keyID     string
	algorithm string
}

// NewAuthenticator returns a new authenticator
// params:
//	- storage - storage type (file or memory for now)
func NewAuthenticator(storage Storage, now time.Time) (*Authenticator, error) {
	publicKeyLookup := NewKeyFunc(storage.Keys())

	// algorithm is globally defined
	if jwt.GetSigningMethod(algorithm) == nil {
		return nil, errors.Errorf("unknown algorithm %v", algorithm)
	}

	parser := jwt.Parser{
		ValidMethods: []string{algorithm},
	}

	// Load the current key from the storage engine
	curKey := storage.Current()
	if curKey == nil {
		return nil, errors.New("Missing private key")
	}

	auth := Authenticator{
		privateKey: curKey,
		keyID:      curKey.keyID,
		algorithm:  algorithm,
		kf:         publicKeyLookup,
		parser:     &parser,
	}

	return &auth, nil
}

// GenerateToken generates a new token based on provided claims
func (auth *Authenticator) GenerateToken(claims Claims) (string, error) {
	method := jwt.GetSigningMethod(auth.algorithm)

	// there should be user data and expiration in claims
	token := jwt.NewWithClaims(method, claims)
	token.Header["kid"] = auth.keyID

	str, err := token.SignedString(auth.privateKey.PrivateKey)
	if err != nil {
		return "", errors.Wrap(err, "signing token")
	}

	return str, nil
}

// ParseClaims decodes token to a claims struct
func (auth *Authenticator) ParseClaims(tokenString string) (Claims, error) {

	f := func(t *jwt.Token) (interface{}, error) {
		kid, ok := t.Header["kid"]
		if !ok {
			return nil, errors.New("Missing key id (kid) in token header")
		}

		kidStr, ok := kid.(string)
		if !ok {
			return nil, errors.New("Unable to convert kid to string")
		}

		// remove if this doesn't work
		// method := jwt.GetSigningMethod(auth.algorithm)
		// if _, ok := t.Method.(method); !ok {
		// 	return nil, errors.New("Unexpected signing method")
		// }

		return auth.kf(kidStr)
	}

	var claims Claims
	token, err := auth.parser.ParseWithClaims(tokenString, &claims, f)
	if err != nil {
		return Claims{}, errors.Wrap(err, "parsing token to claims")
	}

	if !token.Valid {
		return Claims{}, errors.New("Invalid token")
	}

	return claims, nil
}
