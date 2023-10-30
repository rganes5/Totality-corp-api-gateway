package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/api/handlers"
)

func RegisterUserRoutes(api *gin.RouterGroup, userHandler handlers.UserHandler) {
	user := api.Group("/user")
	{
		user.GET("/getbyid/:userId", userHandler.GetUserData)
		user.GET("/getbylist", userHandler.GetUserDataList)
	}
}
