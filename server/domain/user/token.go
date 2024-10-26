package user

import (
	"time"

	"github.com/cockroachdb/errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mrngsht/realworld-goa-react/config"
)

type Claims struct {
	jwt.RegisteredClaims

	UserID string `json:"userId"`
}

func IssueToken(userID uuid.UUID, issuedAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			ExpiresAt: jwt.NewNumericDate(issuedAt.Add(1 * time.Hour)),
		},
		UserID: userID.String(),
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
	var claims Claims
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return config.C.PasswordAuthHmacKey, nil
	})
	if err != nil {
		return Token{}, errors.WithStack(err)
	}

	if claims.ExpiresAt == nil {
		return Token{}, errors.Errorf("claim[exp] value is nil")
	}
	if now.After(claims.ExpiresAt.Time) {
		return Token{}, ErrTokenHasExpired
	}

	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return Token{}, errors.Errorf(`claim[userID] value is empty or invalid ("%s"): %v`, claims.UserID, err)
	}

	return Token{UserID: userID}, nil
}
