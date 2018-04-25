package main 

import ( 
	"net/http" 
	"strconv"
	"github.com/gin-gonic/gin" 
) 

func showIndexPage( c *gin.Context ){ 
	articles := getAllArticles() 
	render( c, gin.H {
		"title" : "Home Page",
		"payload" : articles,
	}, "base.html" )
}

func viewArticle( c *gin.Context ){
	if articleId, err := strconv.Atoi( c.Param( "id" ) ); err == nil {
		if article, err := getArticleById( articleId ); err == nil {
			c.HTML(
				http.StatusOK,
				"article.html",
				gin.H {
					"title" 	: article.Title,
					"payload" 	: article,
				},
			)
		} else {
			c.AbortWithError( http.StatusNotFound, err )
		}
	} else {
		c.AbortWithError( http.StatusNotFound, err )
	}
}