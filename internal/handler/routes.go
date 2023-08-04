package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.createUser)
		auth.POST("/login", h.loginUser)
	}

	api := router.Group("/api")
	api.Use(h.authMiddleware())
	{
		posts := api.Group("/posts")
		{
			posts.GET("/", h.getAllPosts)
			posts.POST("/", h.createPost)
			posts.GET("/:id", h.getPostByID)
			posts.Use(h.checkPermissions())
			posts.DELETE("/:id", h.deletePost)
			posts.PUT("/:id", h.updatePost)
		}
	}

	return router
}
