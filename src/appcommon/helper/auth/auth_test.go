package auth

import (
	"appcommon/appmodel/usermodel"
	userModel "appcommon/appmodel/usermodel"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	user := userModel.UserFullDto{
		CompanyId: 1,
		GroupId:   1,
		UserBasicInfo: usermodel.UserBaseDto{
			Id:    1,
			Name:  "ss",
			Email: "abc@test.com",
		},
	}
	GenerateToken("acc", &user)
	// TODO: Add parse token function to parse the generated token
}
