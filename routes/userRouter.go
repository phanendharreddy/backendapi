package routes

import (
	controller "github.com/HousewareHQ/houseware---backend-engineering-octernship-phanendharreddy/controllers"
	"github.com/HousewareHQ/houseware---backend-engineering-octernship-phanendharreddy/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.DELETE("/users/:user_id", controller.DeleteUser())
}
