package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
)

// Storage provides the ability to persist keys to custom locations.
type Storage interface {
	// Keys returns a map of private keys by kID.
	Keys() map[string]*PrivateKey
	// Current returns the most recently generated private key.
	Current() *PrivateKey
}

// StorageMemory is a storage engine that stores a single private key in memory.
type StorageMemory struct {
	privateKey *PrivateKey
}

// Keys returns a map of private keys by kID.
func (s *StorageMemory) Keys() map[string]*PrivateKey {
	if s == nil || s.privateKey == nil {
		return map[string]*PrivateKey{}
	}
	return map[string]*PrivateKey{
		s.privateKey.keyID: s.privateKey,
	}
}

// Current returns the most recently generated private key.
func (s *StorageMemory) Current() *PrivateKey {
	if s == nil {
		return nil
	}
	return s.privateKey
}

// NewAuthenticatorMemory is a help function that inits a new Authenticator with a single key stored in memory.
func NewAuthenticatorMemory(now time.Time) (*Authenticator, error) {
	storage, err := NewStorageMemory()
	if err != nil {
		return nil, err
	}

	return NewAuthenticator(storage, now)
}

// NewStorageMemory implements the interface Storage to store a single key in memory.
func NewStorageMemory() (*StorageMemory, error) {

	privateKey, err := KeyGen()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate new private key")
	}

	pk, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "parsing auth private key")
	}

	storage := &StorageMemory{
		privateKey: &PrivateKey{
			PrivateKey: pk,
			keyID:      uuid.NewRandom().String(),
			algorithm:  algorithm,
		},
	}

	return storage, nil
}
