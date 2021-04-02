package usermodel

type User struct {
	UserId   int
	UserName string
}

func CreateUser() User {
	return User{UserId: 1, UserName: "qbressler"}
}

func Test() string {
	s := "this is a test"
	return s
}
