package user

import (
	"deep/mock/person"
)

type User struct {
	m person.Male
}

func NewUser(p person.Male) *User {
	return &User{
		m: p,
	}
}

func (u *User) GetUserInfo() error {
	return u.m.Add()
}
