package app

import (
	"github.com/loxt/bookstore-users-api/controllers/ping_controller"
	"github.com/loxt/bookstore-users-api/controllers/users_controller"
)

func mapUrls() {
	router.GET("/ping", ping_controller.Ping)

	router.GET("/users/:user_id", users_controller.GetUser)
	router.POST("/users", users_controller.CreateUser)
}
