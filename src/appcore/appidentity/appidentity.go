// Package appidentity provides methods to manage the unique ID for the server
package appidentity

import (
	"errors"
)

var numArray = [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var alphArray = [26]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
var alphZArray = [26]byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

const identityPrefix = "app_"

// IAppIdentity interface is used to manage the identity string of the running application
type IAppIdentity interface {
	CreateAppIdentity() error
	GetAppIdentity() string
	UpdateAppIdentity() (string, error)
}

// AppIdentity struct contains the identityFlag to indicate the position of the running server,
// and the IdentityString for its unique id
type AppIdentity struct {
	identityFlag int
	// {prefix}
	identityString string // TODO: Define the structure of the IdentityString
}

// CreateAppIdentity function creates the identity string for the application
func (app *AppIdentity) CreateAppIdentity() error {
	// TODO: Need to combine random number and current system time

	//now := time.Now()
	//unixTime := now.Unix()
	nearBys := findNearBy()
	app.identityFlag = len(nearBys) + 1
	app.identityString = "temp-app_identity"

	return nil
}

// GetAppIdentity function returns the identity string of the application
func (app *AppIdentity) GetAppIdentity() string {
	return app.identityString
}

// UpdateAppIdentity function generates a new identity string for the string
func (app *AppIdentity) UpdateAppIdentity() {
	// TODO: Implement the same logic as CreateAppIdentity function
}

// TODO: Create method to generate random string

// findNearBy function will find connected server
func findNearBy() []string {
	return make([]string, 0)
}

// select number by using the provided index
func selectNum(index int) ([]byte, error) {
	if index < 0 {
		err := errors.New("Invalid index to select number")
		return nil, err
	}
	return numArray[index:index], nil
}

// select alpha by using the provided index
func selectAlpha(index int) ([]byte, error) {
	if index < 0 {
		err := errors.New("Invalid index to select alpha")
		return nil, err
	}
	return alphArray[index:index], nil
}

// select captial alpha by using the provided index
func selectCaptialAlpha(index int) ([]byte, error) {
	if index < 0 {
		err := errors.New("Invalid index to select capital alpha")
		return nil, err
	}
	return alphZArray[index:index], nil
}
