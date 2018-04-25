package main

func initializeRoutes(){
	router.GET( "/", showIndexPage )
	
	router.GET( "/article/view/:id", viewArticle )

	userRoutes := router.Group( "/v1/user" )
	{
		userRoutes.GET( "/register", ensureNotLoggedIn(), showRegistrationPage )
		userRoutes.POST( "/register", ensureNotLoggedIn(), register )
		userRoutes.GET( "/login", ensureNotLoggedIn(), showLoginPage )
		userRoutes.POST( "/login", ensureNotLoggedIn(), performLogin )
		userRoutes.GET( "/logout", ensureLoggedIn(), logout )
	}
}