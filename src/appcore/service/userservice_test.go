package service

import (
	userModel "appcommon/appmodel/usermodel"
	"testing"
)

func TestCreateNewUserShouldSuccess(t *testing.T) {
	// setup user service for testing
	testUserService := UserService{}
	testUserService.Setup()
	defer testUserService.Repository.CloseMe()
	// test user data
	newUser := userModel.UserFullDto{
		UserBasicInfo: userModel.UserBaseDto{
			Id:       1,
			Email:    "testuser@test.com",
			Name:     "test new user",
			Password: "password",
		},
		GroupId:   1,
		CompanyId: 1,
	}
	err := testUserService.CreateNewUser(&newUser)
	if err != nil {
		t.Errorf("New user creation testing failed: %s\n", err)
	}
}
