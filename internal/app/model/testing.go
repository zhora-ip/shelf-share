package model

import "testing"

func TestUser(t *testing.T) *User {
	t.Helper()
	return &User{
		Nickname: "Vasiliy",
		Password: "1231231",
		Email:    "vasyliy@mail.ru",
	}
}

func TestUsers(t *testing.T) []*User {
	t.Helper()
	return []*User{
		{
			Nickname: "Vasiliy",
			Password: "password",
			Email:    "vasyliy@mail.ru",
		},
		{
			Nickname: "Andrey",
			Password: "password",
			Email:    "Andrey@mail.ru",
		},
	}
}
