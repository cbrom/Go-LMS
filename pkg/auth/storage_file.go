package auth

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
)

// StorageFile is a storage engine that stores private keys on the local file system.
type StorageFile struct {
	// Local directory for storing private keys.
	localDir string
	// Duration for keys to be valid.
	keyExpiration time.Duration
	// Map of keys by kid (version id).
	keys map[string]*PrivateKey
	// The current active key to be used.
	curPrivateKey *PrivateKey
}

// Keys returns a map of private keys by kID.
func (s *StorageFile) Keys() map[string]*PrivateKey {
	if s == nil || s.keys == nil {
		return map[string]*PrivateKey{}
	}
	return s.keys
}

// Current returns the most recently generated private key.
func (s *StorageFile) Current() *PrivateKey {
	if s == nil {
		return nil
	}
	return s.curPrivateKey
}

// NewAuthenticatorFile is a help function that inits a new Authenticator
// using the file storage.
func NewAuthenticatorFile(localDir string, now time.Time, keyExpiration time.Duration) (*Authenticator, error) {
	storage, err := NewStorageFile(localDir, now, keyExpiration)
	if err != nil {
		return nil, err
	}

	return NewAuthenticator(storage, now)
}

// NewStorageFile implements the interface Storage to support persisting private keys
// to the local file system.
func NewStorageFile(localDir string, now time.Time, keyExpiration time.Duration) (*StorageFile, error) {
	if localDir == "" {
		localDir = filepath.Join(os.TempDir(), "auth-private-keys")
	}

	if _, err := os.Stat(localDir); os.IsNotExist(err) {
		err = os.MkdirAll(localDir, os.ModePerm)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create storage directory %s", localDir)
		}
	}

	storage := &StorageFile{
		localDir:      localDir,
		keyExpiration: keyExpiration,
		keys:          make(map[string]*PrivateKey),
	}

	if now.IsZero() {
		now = time.Now().UTC()
	}

	// Time threshold to stop loading keys, any key with a created date
	// before this value will not be loaded.
	var disabledCreatedDate time.Time

	// Time threshold to create a new key. If a current key exists and the
	// created date of the key is before this value, a new key will be created.
	var activeCreatedDate time.Time

	// If an expiration duration is included, convert to past time from now.
	if keyExpiration.Seconds() != 0 {
		// Ensure the expiration is a time in the past for comparison below.
		if keyExpiration.Seconds() > 0 {
			keyExpiration = keyExpiration * -1
		}
		// Stop loading keys when the created date exceeds two times the key expiration
		disabledCreatedDate = now.UTC().Add(keyExpiration * 2)

		// Time used to determine when a new key should be created.
		activeCreatedDate = now.UTC().Add(keyExpiration)
	}

	// Values used to format filename.
	filePrefix := "sassauth_"
	fileExt := ".privatekey"

	files, err := ioutil.ReadDir(localDir)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to list files in directory %s", localDir)
	}

	// Map of keys stored by version id. version id is kid.
	keyContents := make(map[string][]byte)

	// The current key id if there is an active one.
	var curKeyID string

	// The max created data to determine the most recent key.
	var lastCreatedDate time.Time

	for _, f := range files {
		if !strings.HasPrefix(f.Name(), filePrefix) || !strings.HasSuffix(f.Name(), fileExt) {
			continue
		}

		// Extract the created timestamp and kID from the filename.
		fname := strings.TrimSuffix(f.Name(), fileExt)
		pts := strings.Split(fname, "_")
		if len(pts) != 3 {
			return nil, errors.Errorf("unable to parse filename %s", f.Name())
		}
		createdAt := pts[1]
		kID := pts[2]

		// Covert string timestamp to int.
		createdAtSecs, err := strconv.Atoi(createdAt)
		if err != nil {
			return nil, errors.Wrapf(err, "failed parse timestamp from %s", f.Name())
		}
		ts := time.Unix(int64(createdAtSecs), 0)

		// If the created time of the key is less than the disabled threshold, skip.
		if !disabledCreatedDate.IsZero() && ts.UTC().Unix() < disabledCreatedDate.UTC().Unix() {
			continue
		}

		filePath := filepath.Join(localDir, f.Name())
		dat, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, errors.Wrapf(err, "failed read file %s", f.Name())
		}

		keyContents[kID] = dat

		if lastCreatedDate.IsZero() || ts.UTC().Unix() > lastCreatedDate.UTC().Unix() {
			curKeyID = kID
			lastCreatedDate = ts.UTC()
		}
	}

	//
	if !activeCreatedDate.IsZero() && lastCreatedDate.UTC().Unix() < activeCreatedDate.UTC().Unix() {
		curKeyID = ""
	}

	// If there are no keys or the current key needs to be rotated, generate a new key.
	if len(keyContents) == 0 || curKeyID == "" {
		privateKey, err := KeyGen()
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate new private key")
		}

		kID := uuid.NewRandom().String()

		fname := fmt.Sprintf("%s%d_%s%s", filePrefix, now.UTC().Unix(), kID, fileExt)

		filePath := filepath.Join(localDir, fname)

		err = ioutil.WriteFile(filePath, privateKey, 0644)
		if err != nil {
			return nil, errors.Wrapf(err, "failed write file %s", filePath)
		}

		keyContents[curKeyID] = privateKey
	}

	// Loop through all the key bytes and load the private key.
	for kid, key := range keyContents {
		pk, err := jwt.ParseRSAPrivateKeyFromPEM(key)
		if err != nil {
			return nil, errors.Wrap(err, "parsing auth private key")
		}

		storage.keys[kid] = &PrivateKey{
			PrivateKey: pk,
			keyID:      kid,
			algorithm:  algorithm,
		}

		if kid == curKeyID {
			storage.curPrivateKey = storage.keys[kid]
		}
	}

	return storage, nil
}
