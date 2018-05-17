package main

import (
	"strings"
	//	"math/rand"
	"net/http"
	//    "strconv"
	"github.com/gin-gonic/gin"
)

/*func generateSessionToken() string {
	return strconv.FormatInt( rand.Int63(), 16 )
}

func showRegistrationPage( c *gin.Context ){
	render(
		c,
		gin.H {
			"title" : "Register",
		},
		"register.html",
 	)
}

func register( c *gin.Context ){
	username := c.PostForm( "username" )
	password := c.PostForm( "password" )

	if _, err := registerUser( username, password ); err == nil {
		token := generateSessionToken()
		c.SetCookie( "token", token, 3600, "", "", false, true )
		c.Set( "is_logged_in", true )

		render(
			c,
			gin.H {
				"title" : "Success",
			},
			"login-success.html",
		)
	} else {
		c.HTML(
			http.StatusBadRequest,
			"register.html",
			gin.H {
				"errorTitle" : "Registration failed",
				"errorMessage" : err.Error(),
			},
		)
	}
}

func showLoginPage( c *gin.Context ){
	render(
		c,
		gin.H {
			"title" : "Login",
		},
		"login.html",
	)
}

func performLogin( c *gin.Context ){
	username := c.PostForm( "username" )
	password := c.PostForm( "password" )

	if isUserValid( username, password ) {
		token := generateSessionToken()
        c.SetCookie( "token", token, 3600, "", "", false, true )

		render(
			c,
			gin.H {
				"title": "Successful Login",
			},
			"login-success.html",
		)
	} else {
		c.HTML(
			http.StatusBadRequest,
			"login.html",
			gin.H {
				"errorTitle" : "Login Failed",
				"errorMessage" : "Invalid credentials provided",
			},
		)
	}
}

func logout( c *gin.Context ){
	c.SetCookie( "token", "", -1, "", "", false, true )
    c.Redirect( http.StatusTemporaryRedirect, "/" )
}*/

func userRegisterAction(c *gin.Context) {
	type Login struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var json Login

	if c.BindJSON(&json) == nil {
		if user, err := createUser(env.db, json.Email, json.Password); err == nil {
			c.JSON(http.StatusCreated, gin.H{
				"user": user,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
	}
}

func userLoginAction(c *gin.Context) {
	type Login struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var json Login

	if c.BindJSON(&json) == nil {
		if user, err := loginUser(env.db, json.Email, json.Password); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user": user,
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
	}
}

func userLogoutAction(c *gin.Context) {
	token := strings.Replace(c.Request.Header.Get("Authorization"), "token=", "", 1)
	status := logoutUser(env.db, token)
	if !status {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong",
		})
	}

	c.Writer.WriteHeader(http.StatusOK)
}
