package router

import (
	"ParkNavigate/global"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(global.Config.Server.Env)
	router := gin.Default()
	apiGroup := router.Group("/api/")

	// Navigation service
	navApiGroup := apiGroup.Group("/navigation/")
	//navApiGroup.POST("login", http.UserLogin)

	return router
}
