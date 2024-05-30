package routes

import (
	"educative-1/controllers"
	middleware "educative-1/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router gin.Engine) {
	router.Use(middleware.AuthenticateUser())
	router.GET("/users/:user_id", controllers.GetUser())
}
