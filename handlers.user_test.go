package main

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"net/http/httptest"
	"strings"
	"strconv"
	"testing"
)

func getLoginPostPayload() string {
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
}