package route

import (
	"net/http"

	"github.com/errntry/errntry/internal/api/controller/user"
	"github.com/errntry/errntry/internal/config"
	"github.com/gin-gonic/gin"
)

func Run() error {
	cfg := config.Load()

	if appEnv := cfg.Get("APP_ENV"); appEnv != "" {
		switch appEnv {
		case "debug", "test":
			gin.SetMode(gin.DebugMode)
		case "release", "production":
			gin.SetMode(gin.ReleaseMode)
		}
	}

	router := gin.Default()
	api := router.Group("/api/v1/")
	{
		api.GET("/users", user.Index)
		// TODO: Create should be removed after creating a proper signup endpoint
		api.POST("/users", user.Create)
		api.GET("/users/:id", user.Get)
		api.PUT("/users/:id", user.Update)
		api.DELETE("/users/:id", user.Delete)
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	return router.Run(":" + cfg.Get("SERVER_PORT"))
}
