package user

import (
	"github.com/cockroachdb/errors"
	"golang.org/x/crypto/bcrypt"
)

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
