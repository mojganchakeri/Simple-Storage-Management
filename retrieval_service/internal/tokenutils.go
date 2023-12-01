package internal

import (
	"fmt"
	"retreival_service/internal/models"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var whiteListTokens []string

// splict jwt from authHeader
func ExtractAuthToken(authHeader string) string {
	t := strings.Split(authHeader, " ")
	if len(t) == 2 {
		return t[1]
	}
	return ""
}

func CreateAccessToken(user models.User, secret string, expiry int) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claims := models.JwtCustomClaims{
		Name: user.Username,
		ID:   user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func CreateRefreshToken(user models.User, secret string, expiry int) (refreshToken string, err error) {
	claims := models.JwtCustomRefreshClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expiry)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("Invalid Token")
	}

	return claims["id"].(string), nil
}

func AddAccessTokenToWhiteList(accessToken string) {
	whiteListTokens = append(whiteListTokens, accessToken)
}

func RemoveAccessTokenFromWhiteList(authHeader string) {
	authToken := ExtractAuthToken(authHeader)
	if authToken != "" {
		whiteListTokens = remove(whiteListTokens, authToken)
	}
}

func IsWhitelisted(accessToken string) bool {
	for _, v := range whiteListTokens {
		if v == accessToken {
			return true
		}
	}
	return false
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
