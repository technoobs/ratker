package authmodel

import (
	"github.com/dgrijalva/jwt-go"
)

// TODO: Need to assign JSON property name

// ** Will be replaced by Protobuf **
// UserJwt is the structure used for JWT token
type UserJwt struct {
	Id          int16  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	StandardJwt jwt.StandardClaims
}

// Valid function validates user's jwt token
func (c UserJwt) Valid() error {
	stdErr := c.StandardJwt.Valid()
	if stdErr != nil {
		return stdErr
	}
	return nil
}

func (c UserJwt) validatesUser() error {
	return nil
}
