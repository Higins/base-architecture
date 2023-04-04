package handlers

import (
	"fmt"
	"time"

	"github.com/Higins/blocks/constants"
	"github.com/golang-jwt/jwt/v4"
)

const exp = time.Minute * 10

// Create a JWT token signed with jwtSecret, containing claims. Claims should be primitive string or int values.
func NewSignedTokenWithClaims(jwtSecret string, claims map[string]interface{}, expiration ...time.Duration) (string, error) {
	if len(expiration) == 1 {
		claims["exp"] = time.Now().UTC().Add(expiration[0]).Unix()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))

	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Verify that the signedToken JWT token was signed with jwtSecret. Returns claims of the JWT token.
func (a *AuthHandler) DecodeClaims(jwtSecret string, signedToken string, needExpiration bool) (map[string]interface{}, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})
	if err != nil {
		return map[string]interface{}{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if needExpiration {
			if _, ok := claims["exp"]; !ok {
				return map[string]interface{}{}, constants.ErrTokenExpired
			}
		}

		return claims, nil
	} else {
		return map[string]interface{}{}, err
	}
}
