package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth") 
	{
		auth.POST("/register", h.createUser)
		auth.POST("/login", h.loginUser)
	}

	return router
}
