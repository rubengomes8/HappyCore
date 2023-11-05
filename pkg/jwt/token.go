package auth

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt"
)

type TokenService struct {
	apiSecret          string
	tokenLifespanHours int
}

func NewTokenService(
	apiSecret string,
	tokenLifespanHours int,
) TokenService {
	return TokenService{
		apiSecret:          apiSecret,
		tokenLifespanHours: tokenLifespanHours,
	}
}

func (s TokenService) GenerateToken(sub uint) (string, error) {

	if s.apiSecret == "" {
		return "", ErrAPISecretRequired
	}

	if s.tokenLifespanHours == 0 {
		return "", ErrTokenLifecycleRequired
	}

	claims := jwt.MapClaims{}
	claims["sub"] = sub
	claims["exp"] = time.Now().UTC().Add(time.Hour * time.Duration(s.tokenLifespanHours)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.apiSecret))
}

func (s TokenService) ValidateToken(ctx *gin.Context) error {

	tokenString, err := getBearerToken(ctx)
	if err != nil {
		return errors.New("invalid authentication")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.apiSecret), nil
	})

	if !token.Valid {
		return errors.New("invalid token")
	}

	return err
}

func (s TokenService) ExtractClaimSub(ctx *gin.Context) (int, error) {

	tokenString, err := getBearerToken(ctx)
	if err != nil {
		return 0, errors.New("invalid authentication")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.apiSecret), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	sub, ok := claims["sub"].(float64)
	if !ok {
		return 0, errors.New("invalid claim sub")
	}

	return int(sub), nil
}

func getBearerToken(ctx *gin.Context) (string, error) {
	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) != 2 {
		return "", errors.New("invalid authentication")
	}
	return strings.Split(bearerToken, " ")[1], nil
}
