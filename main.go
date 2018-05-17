package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

//var db *gorm.DB
//var router *gin.Engine

// Env is a structure that keeps app global environment handlers
type Env struct {
	db     DataAccessLayer
	router *gin.Engine
}

var env Env

func main() {
	// Lets set port hardly for now
	os.Setenv("PORT", "9000")

	// Init database connection
	db := NewDB("root:qwas1234@/gr?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	// Set database and router for current application environement
	env.db = db
	env.router = gin.Default()

	// Init routes
	initializeRoutes()

	// Start serving the application
	env.router.Run()
}

/*func render( c *gin.Context, data gin.H, templateName string ){
	switch c.Request.Header.Get("Accept") {
		case "application/json":
			c.JSON( http.StatusOK, data["payload"] )
		case "application/xml":
			c.XML( http.StatusOK, data["payload"] )
		default:
			c.HTML( http.StatusOK, templateName, data )
	}
}*/
