package routes

import (
	"github.com/eranamarante/go-expense-tracker-api/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) {
	route.POST("/users/signup", controllers.SignUp())
	route.POST("/users/login", controllers.Login())
}
