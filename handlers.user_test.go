package main

import (

	/*"io/ioutil"
	"net/url"
	"net/http"
	"net/http/httptest"
	"strings"
	"strconv"*/
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/*func getLoginPostPayload() string {
	params := url.Values{}
	params.Add( "username", "user1" )
	params.Add( "password", "pass1" )

	return params.Encode()
}

func getRegistrationPostPayload() string {
	params := url.Values{}
	params.Add( "username", "u1" )
	params.Add( "password", "p1" )

	return params.Encode()
}

func TestShowregistrationPageUnauthenticated( t *testing.T ){
	r := getRouter( true )
	r.GET( "/user/register", showRegistrationPage )

	req, _ := http.NewRequest( "GET", "/user/register", nil )
	testHTTPResponse( t, r, req, func( w *httptest.ResponseRecorder ) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll( w.Body )
		pageOK := err == nil && strings.Index( string( p ), "Register" ) > 0

		return statusOK && pageOK
	} )
}

func TestRegisterUnauthenticated( t *testing.T ){
	saveLists()

	w := httptest.NewRecorder()
	r := getRouter( true )
	r.POST( "/user/register", register )

	registrationPayload := getRegistrationPostPayload()
	req, _ := http.NewRequest( "POST", "/user/register", strings.NewReader( registrationPayload ) )
	req.Header.Add( "Content-Type", "application/x-www-form-urlencoded" )
	req.Header.Add( "Content-Length", strconv.Itoa( len( registrationPayload ) ) )

	r.ServeHTTP( w, req )

	if w.Code != http.StatusOK {
		t.Fail()
	}

	p, err := ioutil.ReadAll( w.Body )
	if err != nil || strings.Index( string( p ), "Success" ) < 0 {
		t.Fail()
	}

	restoreLists()
}

func TestRegisterUnauthenticatedUnavailableUsername( t *testing.T ){
	saveLists()

	w := httptest.NewRecorder()
	r := getRouter( true )
	r.POST( "/user/register", register )

	registrationPayload := getLoginPostPayload()
	req, _ := http.NewRequest( "POST", "/user/register", strings.NewReader( registrationPayload ) )
	req.Header.Add( "Content-Type", "application/x-www-form-urlencoded" )
	req.Header.Add( "Content-Length", strconv.Itoa( len( registrationPayload ) ) )

	r.ServeHTTP( w, req )

	if w.Code != http.StatusBadRequest {
		t.Fail()
	}

	restoreLists()
}

func TestShowLoginPageUnauthenticated( t *testing.T ){
	r := getRouter( true )
	r.GET( "/user/login", showLoginPage )

	req, _ := http.NewRequest( "GET", "/user/login", nil )
	testHTTPResponse( t, r, req, func( w *httptest.ResponseRecorder ) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll( w.Body )
		pageOK := err == nil && strings.Index( string( p ), "Login" ) > 0

		return statusOK && pageOK
	} )
}

func TestLoginUnauthenticated( t *testing.T ) {
	saveLists()

    w := httptest.NewRecorder()
    r := getRouter( true )

    r.POST( "/user/login", performLogin )

    loginPayload := getLoginPostPayload()
    req, _ := http.NewRequest( "POST", "/user/login", strings.NewReader( loginPayload ) )
    req.Header.Add( "Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add( "Content-Length", strconv.Itoa( len( loginPayload ) ) )

    r.ServeHTTP( w, req )

    if w.Code != http.StatusOK {
        t.Fail()
    }

    p, err := ioutil.ReadAll( w.Body )
    if err != nil || strings.Index( string( p ), "Success" ) < 0 {
        t.Fail()
	}

    restoreLists()
}

func TestLoginUnauthenticatedIncorrectCredentials( t *testing.T ) {
	saveLists()

    w := httptest.NewRecorder()
    r := getRouter( true )

    r.POST( "/user/login", performLogin )

    loginPayload := getRegistrationPostPayload()
    req, _ := http.NewRequest( "POST", "/user/login", strings.NewReader( loginPayload ) )
    req.Header.Add( "Content-Type", "application/x-www-form-urlencoded" )
    req.Header.Add( "Content-Length", strconv.Itoa( len( loginPayload ) ) )

    r.ServeHTTP( w, req )

    if w.Code != http.StatusBadRequest {
        t.Fail()
	}

    restoreLists()
}*/

// Test user register action when user email already in use. Test must return 400 status code and error message
func TestUserRegisterEmailAlreadyExists(t *testing.T) {
	// Mock database calls
	env.db = &mockDatabase{}

	// Create router and serve register call
	r := getRouter()
	r.POST("/v1/user/register", userRegisterAction)

	w := httptest.NewRecorder()
	var jsonPayload = []byte(`{"email":"exists@gentlereader.com","password":"test"}`)
	req, _ := http.NewRequest("POST", "/v1/user/register", bytes.NewBuffer(jsonPayload))
	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fail()
	}

	if strings.Index(w.Body.String(), "Email address already in use") < 0 {
		t.Fail()
	}
}

// Test user register when we pass wrong payload. Test must return 400 status code and error message
func TestUserRegisterWrongPayload(t *testing.T) {
	// Mock database calls
	env.db = &mockDatabase{}

	// Create router and serve register call
	r := getRouter()
	r.POST("/v1/user/register", userRegisterAction)

	w := httptest.NewRecorder()
	jsonPayload := []byte(`{"email":"exists`)
	req, _ := http.NewRequest("POST", "/v1/user/register", bytes.NewBuffer(jsonPayload))
	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fail()
	}

	if strings.Index(w.Body.String(), "Bad request") < 0 {
		t.Fail()
	}
}

// Test user register when we pass correct payload. Test must return 201 status code and created user details
func TestUserRegister(t *testing.T) {
	// Mock database calls
	env.db = &mockDatabase{}

	// Create router and serve register call
	r := getRouter()
	r.POST("/v1/user/register", userRegisterAction)

	w := httptest.NewRecorder()
	jsonPayload := []byte(`{"email":"user1@gentlereader.com","password":"test"}`)
	req, _ := http.NewRequest("POST", "/v1/user/register", bytes.NewBuffer(jsonPayload))
	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fail()
	}

	if strings.Index(w.Body.String(), "user") < 0 {
		t.Fail()
	}
}

// Test user login when pass wrong credentials
func TestUserLoginWrongCredentials(t *testing.T) {
	// Mock database calls
	env.db = &mockDatabase{}

	// Create router and serve register call
	r := getRouter()
	r.POST("/v1/user/login", userLoginAction)

	w := httptest.NewRecorder()
	jsonPayload := []byte(`{"email":"notexists@gentlereader.com","password":"test"}`)
	req, _ := http.NewRequest("POST", "/v1/user/login", bytes.NewBuffer(jsonPayload))
	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code == http.StatusOK {
		t.Fail()
	}
}

// Test user login when pass incorrect payload
func TestUserLoginIncorrectPayload(t *testing.T) {
	// Mock database calls
	env.db = &mockDatabase{}

	// Create router and serve register call
	r := getRouter()
	r.POST("/v1/user/login", userLoginAction)

	w := httptest.NewRecorder()
	jsonPayload := []byte(`{"email":"no`)
	req, _ := http.NewRequest("POST", "/v1/user/login", bytes.NewBuffer(jsonPayload))
	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code == http.StatusOK {
		t.Fail()
	}
}

// Test user login when pass correct payload
func TestUserLogin(t *testing.T) {
	// Mock database calls
	env.db = &mockDatabase{}

	// Create router and serve register call
	r := getRouter()
	r.POST("/v1/user/login", userLoginAction)

	w := httptest.NewRecorder()
	jsonPayload := []byte(`{"email":"exists@gentlereader.com","password":"test"}`)
	req, _ := http.NewRequest("POST", "/v1/user/login", bytes.NewBuffer(jsonPayload))
	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	if strings.Index(w.Body.String(), "user") < 0 {
		t.Fail()
	}
}

// Test user logout with failed response
func TestUserLogoutFailed(t *testing.T) {
	// Mock database calls
	env.db = &mockDatabase{}

	// Create router and serve register call
	r := getRouter()
	r.GET("/v1/user/logout", userLogoutAction)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/user/logout", nil)
	req.Header.Add("Authorization", "token=invalid")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

// Test user logout with success response
func TestUserLogout(t *testing.T) {
	// Mock database calls
	env.db = &mockDatabase{}

	// Create router and serve register call
	r := getRouter()
	r.GET("/v1/user/logout", userLogoutAction)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/user/logout", nil)
	req.Header.Add("Authorization", "token=valid")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}
}
