package domain

type User struct {
	name string
	age  int32
}

func NewUser(name string, age int32) *User {
	return &User{name: name, age: age}
}