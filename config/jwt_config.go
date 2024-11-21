package config

import (
	"slices"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Login string `json:"login"`
	jwt.RegisteredClaims
}

func GetJWTConfig() echojwt.Config {

	return echojwt.Config{
		NewClaimsFunc: func(e echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("manoel"),
		Skipper: func(e echo.Context) bool {
			return slices.Contains(
				[]string{"/user", "/auth"},
				e.Request().URL.Path,
			)
		},
	}
}

func GenerateToken(login string) (string, error) {
	claims := jwtCustomClaims{
		login,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("manoel"))
	if err != nil {
		return "", err
	}

	return t, nil
}

func GetTokenInfo(token string) (string, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("manoel"), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return claims["login"].(string), nil
	}
	return "", nil
}
