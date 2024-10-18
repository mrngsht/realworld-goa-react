package user

import (
	"crypto/rand"
	"log"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mrngsht/realworld-goa-react/config"
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

func IssueToken(userID uuid.UUID, issuedAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": issuedAt.Unix(),
		"exp": issuedAt.Add(1 * time.Hour).Unix(),
		"uid": userID.String(),
	})
	tokenString, err := token.SignedString(config.C.PasswordAuthHmacKey)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return tokenString, nil
}

func VerifyToken(tokenString string, now time.Time) (userID uuid.UUID, ok bool) {
	userID, err := func() (uuid.UUID, error) {
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return config.C.PasswordAuthHmacKey, nil
		})
		if err != nil {
			return uuid.Nil, errors.WithStack(err)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return uuid.Nil, errors.Errorf("failed to cast token.Claims to jwt.MapClaims")
		}

		exp, ok := claims["exp"].(int64)
		if !ok {
			return uuid.Nil, errors.Errorf("claim[exp] value is invalid: %v\n", claims["exp"])
		}
		expt := time.Unix(exp, 0)
		if now.After(expt) {
			return uuid.Nil, errors.Errorf("token has expired at: %v\n", expt)
		}

		uid, ok := claims["uid"].(string)
		if !ok {
			return uuid.Nil, errors.Errorf("claim[uid] value is invalid: %v\n", claims["uid"])
		}

		uUUID, err := uuid.Parse(uid)
		if !ok {
			return uuid.Nil, errors.WithStack(err)
		}

		return uUUID, nil
	}()
	if err != nil {
		log.Default().Printf("error during veifying token: %v", err)
		return uuid.Nil, false
	}

	return userID, true
}
