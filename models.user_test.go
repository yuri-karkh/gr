package main

import (
	"testing"
)

/*func TestCheckUserEmailIsAvailable(t *testing.T) {
	mdb := &mockDatabase{}

	emailAvailable := mdb.CheckUserEmailIsAvailable("exists@gentlereader.com")
	if emailAvailable {
		t.Fail()
	}

	emailAvailable = mdb.CheckUserEmailIsAvailable("notexists@gentlereader.com")
	if !emailAvailable {
		t.Fail()
	}
}*/

func TestCreateUser(t *testing.T) {
	mdb := &mockDatabase{}

	user, err := createUser(mdb, "exists@gentlereader.com", "test")
	if user.ID != 0 || err == nil {
		t.Fail()
	}

	user, err = createUser(mdb, "user1@gentlereader.com", "test")
	if user.ID != 1 || err != nil {
		t.Fail()
	}

	user, err = createUser(mdb, "user2@gentlereader.com", "test")
	if user.ID != 0 || err == nil {
		t.Fail()
	}
}

func TestLoginUser(t *testing.T) {
	mdb := &mockDatabase{}

	user, err := loginUser(mdb, "notexists@gentlereader.com", "test")
	if user.ID != 0 || err == nil {
		t.Fail()
	}

	user, err = loginUser(mdb, "exists@gentlereader.com", "test")
	if user.ID != 1 || err != nil {
		t.Fail()
	}
}

func TestLogoutUser(t *testing.T) {
	mdb := &mockDatabase{}

	status := logoutUser(mdb, "invalid")
	if status {
		t.Fail()
	}

	status = logoutUser(mdb, "valid")
	if !status {
		t.Fail()
	}
}
