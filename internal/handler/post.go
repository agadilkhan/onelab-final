package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/agadilkhan/onelab-final/api"
	"github.com/agadilkhan/onelab-final/internal/entity"
	"github.com/gin-gonic/gin"
)

// @Summary Get All Posts
// @Tags posts
// @Description get all posts
// @ID get all posts
// @Accept json
// @Produce json
// @Success 200 {object} api.Ok
// @Failure 400,404 {object} api.Error
// @Failure 500 {object} api.Error
// @Failure default {object} api.Error
// @Router /api/posts [get]
func (h *Handler) getAllPosts(ctx *gin.Context) {
	posts, err := h.service.GetAllPosts(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.Ok{
		Code:    1,
		Message: "Ok",
		Data:    posts,
	})

	return
}

// @Summary Create Post
// @Tags posts
// @Description create post
// @ID create post
// @Accept json
// @Produce json
// @Param input body api.PostRequest true "post info"
// @Success 200 {object} api.Ok
// @Failure 400,404 {object} api.Error
// @Failure 500 {object} api.Error
// @Failure default {object} api.Error
// @Router /api/posts [post]
func (h *Handler) createPost(ctx *gin.Context) {
	var req api.PostRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
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

	p := &entity.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}

	err = h.service.CreatePost(ctx, p)
	if err != nil {
		log.Printf("can't create post: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "can't create post",
		})
	}

	ctx.JSON(http.StatusCreated, &api.Ok{
		Code: 1,
		Message: "created: true",
		Data: "",
	})
}

// @Summary Get Post by ID
// @Tags posts
// @Description get post by id
// @ID get post by id
// @Accept json
// @Produce json
// @Success 200 {object} api.Ok
// @Failure 400,404 {object} api.Error
// @Failure 500 {object} api.Error
// @Failure default {object} api.Error
// @Router /api/posts/:id [get]
func (h *Handler) getPostByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "can't convert",
		})
		return
	}

	post, err := h.service.GetPostByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -2,
			Message: "not found post",
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.Ok{
		Code:    1,
		Message: "Ok",
		Data:    post,
	})

	return
}

// @Summary Delete Post
// @Tags posts
// @Description delete post
// @ID delete post
// @Accept json
// @Produce json
// @Success 200 {object} api.Ok
// @Failure 400,404 {object} api.Error
// @Failure 500 {object} api.Error
// @Failure default {object} api.Error
// @Router /api/posts/:id [delete]
func (h *Handler) deletePost(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "can't convert",
		})
		return
	}

	err = h.service.DeletePost(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -2,
			Message: "can't delete",
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.Ok{
		Code:    1,
		Message: "deleted: true",
		Data:    "",
	})

	return
}

// @Summary Update Post
// @Tags posts
// @Description update post
// @ID update post
// @Accept json
// @Produce json
// @Param input body api.PostRequest true "post info"
// @Success 200 {object} api.Ok
// @Failure 400,404 {object} api.Error
// @Failure 500 {object} api.Error
// @Failure default {object} api.Error
// @Router /api/posts/:id [put]
func (h *Handler) updatePost(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "can't convert",
		})
	}

	post, err := h.service.GetPostByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -2,
			Message: "not found post",
		})
		return
	}

	var req *api.PostRequest
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -3,
			Message: err.Error(),
		})
		return
	}
	post.Title, post.Content = req.Title, req.Content

	p, err := h.service.UpdatePost(ctx, post)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -4,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.Ok{
		Code:    1,
		Message: "updated: true",
		Data:    p,
	})

	return
}
