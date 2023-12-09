package routers

import (
	"ParkNavigate/app/gateway/internal/http"
	"ParkNavigate/global"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(global.Config.Server.Env)
	router := gin.Default()

	// 配置 CORS
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // 或者您可以指定特定的源，例如：[]string{"http://example-frontend.com"}
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 定义其他路由
	apiGroup := router.Group("/api/")

	// Navigation service
	navApiGroup := apiGroup.Group("/navigation/")
	navApiGroup.POST("parkings", http.ShowParkings)

	return router
}
