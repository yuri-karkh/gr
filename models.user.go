package main

import (
	"errors"
	"strings"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

var userList = []user {
	user { Username : "user1", Password : "pass1" },
	user { Username : "user2", Password : "pass2" },
	user { Username : "user3", Password : "pass3" },
}

func registerUser( username string, password string ) ( *user, error ) {
	if strings.TrimSpace( password ) == "" {
		return nil, errors.New( "Password can't be empty" )
	} else if !isUsernameAvailable( username ) {
		return nil, errors.New( "Username is not available" )
	}

	user := user{ Username : username, Password : password }
	userList = append( userList, user )

	return &user, nil
}

func isUserValid( username string, password string ) bool {
    for _, user := range userList {
        if user.Username == username && user.Password == password {
            return true
        }
    }
    return false
}

func isUsernameAvailable( username string ) bool {
	for _, user := range userList {
		if user.Username == username {
			return false
		}
	}
	return true
}