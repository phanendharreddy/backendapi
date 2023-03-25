package routes

import (
	controller "github.com/HousewareHQ/houseware---backend-engineering-octernship-phanendharreddy/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/createuser", controller.CreateUser())
	incomingRoutes.POST("users/login", controller.Login())
	incomingRoutes.POST("users/logout", controller.Logout())
}
