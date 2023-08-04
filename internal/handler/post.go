package handler

import (
	"log"
	"net/http"
	"fmt"

	"github.com/agadilkhan/onelab-final/api"
	"github.com/agadilkhan/onelab-final/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllPosts(ctx *gin.Context) {

}

func (h *Handler) createPost(ctx *gin.Context) {
	var req api.CreatePostRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	userID, err := getAuthUserID(ctx)
	if err != nil {
		log.Printf("can't get userID")
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "can't get user id from auth",
		})
		return
	}
	fmt.Println(userID)

	p := &entity.Post {
		Title: req.Title,
		Content: req.Content,
		UserID: userID,
	}

	err = h.service.CreatePost(ctx, p)
	if err != nil {
		log.Printf("can't create post: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "can't create post",
		})
	}
}
