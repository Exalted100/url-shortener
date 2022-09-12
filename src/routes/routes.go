package routes

import "github.com/gin-gonic/gin"

func AddAppRoutes(r *gin.Engine) {
	r.GET("/", healthCheck)
	r.POST("/shorten", getShortURLHandler)
	r.GET("/:url", redirectIfURLFound)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
}
