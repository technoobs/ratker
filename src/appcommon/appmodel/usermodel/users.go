package usermodel

// UserBaseDto is the base type for user data structure
type UserBaseDto struct {
	Id       int16
	Email    string
	Name     string
	Password string
}

// UserFullDto is the full type for user data structure
type UserFullDto struct {
	UserBasicInfo UserBaseDto
	GroupId       int16
	CompanyId     int
}

// ** Will be replaced by Protobuf **
// UserLogin is the type for user login
type UserLogin struct {
	Name     string
	Email    string
	Password string
}
