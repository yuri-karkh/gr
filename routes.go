package main

func initializeRoutes() {
	v1 := env.router.Group("/v1")
	{
		v1.POST("/user/register", userRegisterAction)
		v1.POST("/user/login", userLoginAction)
		v1.GET("/user/logout", userLogoutAction)
	}
}
