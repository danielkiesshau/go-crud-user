package routes

import (
	"educative-1/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("users/signup", controllers.Signup())
}
