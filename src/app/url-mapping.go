package app

import "https://github.com/ZunisAvengers/bookstore_users-api/src/controllers"

func mapUrls() {
	router.GET("/favicon.ico", controllers.GetFavicon)
	router.GET("/ping", controllers.Ping)

	router.POST("/api/users", controllers.CreateUser)
	router.GET("/api/users/:id", controllers.GetUsers)
	// router.GET("/api/users/search", controllers.SearchUser)
}
