package auth

import (
	"fmt"

	authModel "appcommon/appmodel/authmodel"
	userModel "appcommon/appmodel/usermodel"

	"github.com/dgrijalva/jwt-go"
)

// GenerateToken creates JWT token according to provided user credentials
func GenerateToken(signSecret string, userInfo userModel.UserFullDto) {
	// TODO: need to embeed provided standard claims into custom claims
	// https://godoc.org/github.com/dgrijalva/jwt-go#StandardClaims

	userJwt := authModel.UserJwt{}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userJwt)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(signSecret))

	fmt.Println(tokenString, err)
}

// ParseToken parses incoming token into friendly data model
func ParseToken(tokenString string) {

}
