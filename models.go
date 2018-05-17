package main

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Database structure represents structure that keeps DB handler
type Database struct {
	*gorm.DB
}

// NewDB method allows to create new Database
func NewDB(source string) *Database {
	db, err := gorm.Open("mysql", source)
	if err != nil {
		panic("Database connection failed")
	}

	db.AutoMigrate(&User{})

	return &Database{db}
}

// DataAccessLayer data access abstraction layer
type DataAccessLayer interface {
	CheckUserEmailIsAvailable(email string) bool
	CreateNewUser(user *User)
	UserLogin(email string, password string) (User, error)
	UserLogout(token string) bool
}

// CheckUserEmailIsAvailable allows to check is user email already in use
func (db *Database) CheckUserEmailIsAvailable(email string) bool {
	var user User
	db.Where("email = ?", email).First(&user)

	if user.ID == 0 {
		return true
	}

	return false
}

// CreateNewUser method allows to create new user. User struct passed by link
func (db *Database) CreateNewUser(user *User) {
	db.Create(&user)
}

// UserLogin method allows to login user by email and password
func (db *Database) UserLogin(email string, password string) (User, error) {
	var user User
	db.Where("email = ? and password = ?", email, password).First(&user)

	if user.ID == 0 {
		return User{}, errors.New("Login failed")
	}

	user.Token = "Some token hash"
	db.Save(&user)

	return user, nil
}

// UserLogout method allows to log out user by resetting its auth token
func (db *Database) UserLogout(token string) bool {
	var user User
	db.Where("token = ?", token).First(&user)

	if user.ID == 0 {
		return false
	}

	user.Token = ""
	db.Save(&user)

	return true
}
