package usermodel

type User struct {
	UserId   int
	UserName string
}

func CreateUser() User {
	return User{UserId: 1, UserName: "qbressler"}
}
