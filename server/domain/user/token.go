package user

import (
	"time"

	"github.com/cockroachdb/errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mrngsht/realworld-goa-react/config"
)

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

type Token struct {
	UserID uuid.UUID
}

func ParseAndVerifyToken(tokenString string, now time.Time) (Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return config.C.PasswordAuthHmacKey, nil
	})
	if err != nil {
		return Token{}, errors.WithStack(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return Token{}, errors.Errorf("failed to cast token.Claims to jwt.MapClaims")
	}

	exp, ok := claims["exp"].(int64)
	if !ok {
		return Token{}, errors.Errorf("claim[exp] value is invalid: %v\n", claims["exp"])
	}
	expt := time.Unix(exp, 0)
	if now.After(expt) {
		return Token{}, ErrTokenHasExpired
	}

	uid, ok := claims["uid"].(string)
	if !ok {
		return Token{}, errors.Errorf("claim[uid] value is invalid: %v\n", claims["uid"])
	}

	uUUID, err := uuid.Parse(uid)
	if !ok {
		return Token{}, errors.WithStack(err)
	}

	return Token{UserID: uUUID}, nil
}
