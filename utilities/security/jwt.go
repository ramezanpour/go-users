package security

import (
	"fmt"
	"time"

	"github.com/ramezanpour/users/models"

	"github.com/ramezanpour/users/utilities/config"

	"github.com/dgrijalva/jwt-go"
)

// Claims stores information about user auth
type Claims struct {
	Username string `json:"username,omitempty"`
	jwt.StandardClaims
}

// ParseToken takes the token string and returns the clmains data
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.GetConfig().TokenSecret), nil
	})

	if err != nil {
		return nil, err
	}

	fmt.Println(token.Claims)

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {

		return claims, nil
	}
	return nil, fmt.Errorf("Token is not authorized")

}

// CreateToken takes user as parameter and returns the token string accordingly
func CreateToken(user *models.User) (string, error) {
	claims := Claims{
		user.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
			Id:        fmt.Sprint(user.ID),
			Issuer:    "users_module",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.GetConfig().TokenSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
