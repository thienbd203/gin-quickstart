package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Setup(r *gin.Engine, db *gorm.DB) {
	// Global middlewares (sẽ thêm logger, cors sau)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"db":      "connected", // có thể check db.Ping() nếu muốn
			"message": "Gin + GORM + MySQL is up!",
		})
	})

	// API v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Xin chào!",
				"env":     "development", // lấy từ config nếu cần
			})
		})
	}
}