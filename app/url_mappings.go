package app

import (
	"github.com/loxt/bookstore-users-api/controllers/ping_controller"
	"github.com/loxt/bookstore-users-api/controllers/users_controller"
)

func mapUrls() {
	router.GET("/ping", ping_controller.Ping)

	router.POST("/users", users_controller.CreateUser)

	userGroup := router.Group("/users/:user_id")
	userGroup.GET("/", users_controller.GetUser)
	userGroup.PUT("/", users_controller.UpdateUser)
	userGroup.PATCH("/", users_controller.UpdateUser)
}
