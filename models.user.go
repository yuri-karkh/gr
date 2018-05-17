package main

import (
	"errors"
)

// User entity
type User struct {
	ID       uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// createUser method allows to create a new user
func createUser(db DataAccessLayer, email string, password string) (User, error) {
	// Lets chek is user email address is available
	if !db.CheckUserEmailIsAvailable(email) {
		return User{}, errors.New("Email address already in use")
	}

	// Init user struct
	user := User{
		Email:    email,
		Password: password,
	}

	// Try to create user
	db.CreateNewUser(&user)

	// If user was not created lets return error
	if user.ID == 0 {
		return User{}, errors.New("User was not created")
	}

	return user, nil
}

// loginUser method allows to login user
func loginUser(db DataAccessLayer, email string, password string) (User, error) {
	user, err := db.UserLogin(email, password)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

// logoutUser method allows to reset user auth token
func logoutUser(db DataAccessLayer, token string) bool {
	return db.UserLogout(token)
}
