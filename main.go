package main 

import ( 
	"net/http"
	"os"
	"github.com/gin-gonic/gin" 
) 

var router *gin.Engine 

func main() { 
	os.Setenv( "PORT", "9000" )
	// Set the router as the default one provided by Gin 
	router = gin.Default() 
	// Process the templates at the start so that they don't have to be loaded 
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob( "templates/admin/*" ) 
	
	// Init routes
	initializeRoutes()
	
	// Start serving the application 
	router.Run() 
}

func render( c *gin.Context, data gin.H, templateName string ){
	switch c.Request.Header.Get("Accept") {
		case "application/json":
			c.JSON( http.StatusOK, data["payload"] )
		case "application/xml":
			c.XML( http.StatusOK, data["payload"] )
		default:
			c.HTML( http.StatusOK, templateName, data )
	}
}