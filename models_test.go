package main

import (
	"errors"
)

type mockDatabase struct{}

func (mdb *mockDatabase) CheckUserEmailIsAvailable(email string) bool {
	if email == "exists@gentlereader.com" {
		return false
	}

	return true
}

func (mdb *mockDatabase) CreateNewUser(user *User) {
	if user.Email == "user1@gentlereader.com" {
		user.ID = 1
	} else if user.Email == "user2@gentlereader.com" {
		user.ID = 0
	}
}

func (mdb *mockDatabase) UserLogin(email string, password string) (User, error) {
	if email == "exists@gentlereader.com" {
		return User{ID: 1}, nil
	}

	return User{}, errors.New("")
}

func (mdb *mockDatabase) UserLogout(token string) bool {
	if token == "valid" {
		return true
	}
	return false
}
