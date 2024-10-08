package user

import (
	"crypto/rand"

	"github.com/cockroachdb/errors"
	"golang.org/x/crypto/bcrypt"
)

// ref: https://snyk.io/jp/blog/secure-password-hashing-in-go/

func GenPasswordSalt() ([]byte, error) {
	salt := make([]byte, 32)

	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	return salt, nil
}

func GenPasswordHash(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return hash, nil
}

func MatchPassword(hash, password []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hash, password); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}
		return false, errors.WithStack(err)
	}
	return true, nil
}
