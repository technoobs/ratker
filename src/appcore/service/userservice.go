package service

import (
	userModel "appcommon/appmodel/usermodel"
	repository "appcommon/helper/repository"
	"errors"
	"fmt"
)

const (
	// TODO: Move the connection string to config file, and set as environment variable
	connectionDNS  string = "tester:password@tcp(localhost:23306)/RatkerDev"
	repositoryName string = "UserService"
)

// UserService structure holds required resources for the user service
type UserService struct {
	Repository *repository.ServiceRepository
}

// Setup function is used to setup the UserService
func (service *UserService) Setup() {
	repository, err := repository.Repository(repositoryName, connectionDNS)
	if err != nil {
		// TODO: Use log helper for logging functionality
		fmt.Println("Error happened when initiating UserService")
	} else {
		service.Repository = repository
	}
}

// CreateNewUser function creates new user
func (service *UserService) CreateNewUser(user *userModel.UserFullDto) error {
	if user == nil {
		return errors.New("Empty User data")
	}
	// Prepare statement for inserting
	stmtIns, err := service.Repository.Conn.Prepare("INSERT INTO Users (Name, Email, Password, EmailConfirmed) VALUES( ?, ?, ?, ? )")
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	// TODO: Need to handle type conversion for "EmailConfirmed"
	_, err = stmtIns.Exec(user.UserBasicInfo.Name, user.UserBasicInfo.Email, user.UserBasicInfo.Password, 0)
	if err != nil {
		return err
	}
	return nil
}
