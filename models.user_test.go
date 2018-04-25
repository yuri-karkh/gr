package main

import (
	"testing"
)

func TestUsernameAvailability( t *testing.T ){
	saveLists()

	if !isUsernameAvailable( "newuser" ) {
		t.Fail()
	}

	if isUsernameAvailable( "user1" ){
		t.Fail()
	}

	registerUser( "newuser", "newpass" )

	if isUsernameAvailable( "newuser" ){
		t.Fail()
	}

	restoreLists()
}

func TestValidUserRegistration( t *testing.T ){
	saveLists()

	u, err := registerUser( "newuser", "newpass" )

	if err != nil || u.Username == "" {
		t.Fail()
	}

	restoreLists()
}

func TestInvalidUserRegistration( t *testing.T ){
	saveLists()

	u, err := registerUser( "user1", "pass1" )

	if err == nil || u != nil {
		t.Fail()
	}

	u, err = registerUser( "newuser", "" )

	if err == nil || u != nil {
		t.Fail()
	}

	restoreLists()
}

func TestUserValidity( t *testing.T ){
	if !isUserValid( "user1", "pass1" ) {
        t.Fail()
    }

    if isUserValid( "user2", "pass1" ) {
        t.Fail()
    }

    if isUserValid( "user1", "" ) {
        t.Fail()
    }

    if isUserValid( "", "pass1" ) {
        t.Fail()
    }

    if isUserValid( "User1", "pass1" ) {
        t.Fail()
    }
}