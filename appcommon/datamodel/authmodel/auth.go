package authmodel

// ** Will be replaced by Protobuf **
// UserJwt is the structure used for JWT token
type UserJwt struct {
	Id    int16
	Name  string
	Email string
	Iat   string
	Exp   string
}
