package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	m "mindlab/internal/api/middleware"
	"time"
)

func SetupRouter(h *Handler) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")

	{
		api.POST("/login", h.Login.Login)
		api.POST("/logout", m.AuthRequired(h.Login.TokenRepo), h.Login.Logout)

		postGroup := api.Group("/posts")
		{
			postGroup.POST("", m.AuthRequired(h.Login.TokenRepo), m.RoleRequired([]string{"admin", "editor"}), h.Post.Create)
			postGroup.GET("", h.Post.Index)
			postGroup.GET("/:id", h.Post.Show)
			postGroup.PUT("/:id", m.AuthRequired(h.Login.TokenRepo), m.RoleRequired([]string{"admin", "editor"}), h.Post.Update)
			postGroup.DELETE("/:id", m.AuthRequired(h.Login.TokenRepo), m.RoleRequired([]string{"admin"}), h.Post.Delete)
		}
	}

	return r
}
