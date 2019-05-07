package usermodel

// userBaseDto is the base type for user data structure
type userBaseDto struct {
	Id    int16
	Email string
	Name  string
}

// ** Will be replaced by Protobuf **
// UserLogin is the type for user login
type UserLogin struct {
	Name     string
	Email    string
	Password string
}
