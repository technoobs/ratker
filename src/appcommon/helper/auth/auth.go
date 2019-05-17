package auth

import (
	"fmt"

	authModel "appcommon/appmodel/authmodel"
	userModel "appcommon/appmodel/usermodel"
	appIdentity "appcore/appidentity"

	"github.com/dgrijalva/jwt-go"
)

// IAppIdentity interface is used to manage the identity string of the running application
type IAppIdentity interface {
	CreateAppIdentity() error
	GetAppIdentity() string
	UpdateAppIdentity() error
}

// GenerateToken creates JWT token according to provided user credentials
func GenerateToken(signSecret string, userInfo *userModel.UserFullDto) {
	// TODO: need to embeed provided standard claims into custom claims
	// https://godoc.org/github.com/dgrijalva/jwt-go#StandardClaims

	identiyObj := new(appIdentity.AppIdentity)

	userJwt := authModel.UserJwt{
		Id:    userInfo.UserBasicInfo.Id,
		Name:  userInfo.UserBasicInfo.Name,
		Email: userInfo.UserBasicInfo.Email,
		StandardJwt: jwt.StandardClaims{
			Id:     "ss",
			Issuer: generateAppIdentity(identiyObj),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userJwt)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(signSecret))

	fmt.Println(tokenString, err)
}

// ParseToken parses incoming token into friendly data model
func ParseToken(tokenString string) {

}

// generate app identity
func generateAppIdentity(identity IAppIdentity) string {
	if err := identity.CreateAppIdentity(); err != nil {
		return ""
	}
	return identity.GetAppIdentity()
}
