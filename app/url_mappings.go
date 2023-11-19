package app

import (
	"github.com/loxt/bookstore-users-api/controllers/ping_controller"
	"github.com/loxt/bookstore-users-api/controllers/users_controller"
)

func mapUrls() {
	router.GET("/ping", ping_controller.Ping)

	router.POST("/users", users_controller.Create)
	router.POST("/users/login", users_controller.Login)

	router.GET("/internal/users/search", users_controller.Search)

	userGroup := router.Group("/users/:user_id")
	userGroup.GET("/", users_controller.Read)
	userGroup.PUT("/", users_controller.Update)
	userGroup.PATCH("/", users_controller.Update)
	userGroup.DELETE("/", users_controller.Delete)
}
