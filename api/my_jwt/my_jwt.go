package my_jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var Salt = []byte("V3rY_$ecrEt_strInG_ple@Se_donT_hacK")

type Claims struct {
	UID uint  `form:"uid" json:"uid" binding:"required"`
	Exp int64 `form:"exp" json:"exp" binding:"required"`
	// IP    `form:"ip" json:"ip" binding:"required"`
}

func (c *Claims) Valid() error {
	if c.Exp < time.Now().Unix() {
		return errors.New("token expired")
	}
	return nil
}

func ValidateToken(tokenString string) (err error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return Salt, nil
		},
	)
	if err != nil {
		return err
	}

	_, ok := token.Claims.(*Claims)
	if !ok {
		return err
	}

	return nil
}

func ExtractID(tokenStr string) (int, error) {
	if err := ValidateToken(tokenStr); err != nil {
		return 0, err
	}

	claims := &Claims{}
	_, err := jwt.ParseWithClaims(
		tokenStr,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return Salt, nil
		},
	)
	if err != nil {
		return 0, err
	}

	return int(claims.UID), nil
}
