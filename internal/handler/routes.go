package handler

import (
	_ "github.com/agadilkhan/onelab-final/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	// swagger add
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
