package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/agadilkhan/onelab-final/api"
	"github.com/gin-gonic/gin"
)

const authUserID = "userID"

func (h *Handler) authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("authorization")
		if authHeader == "" {
			err := errors.New("authorization header is not set")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &api.Error{
				Code:    -1,
				Message: err.Error(),
			})
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			err := errors.New("authorization header incorrect format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &api.Error{
				Code:    -2,
				Message: err.Error(),
			})
			return
		}

		userID, err := h.service.VerifyToken(fields[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &api.Error{
				Code:    -3,
				Message: err.Error(),
			})
			return
		}

		ctx.Set(authUserID, userID)
		ctx.Next()
	}
}

func (h *Handler) checkPermissions() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		current_user, _ := getAuthUserID(ctx)
		postId, _ := strconv.Atoi(ctx.Param("id"))

		post, err := h.service.GetPostByID(ctx, postId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, &api.Error{
				Code: -1,
				Message: err.Error(),
			})
			return
		}

		if current_user != post.UserID {
			ctx.AbortWithStatusJSON(http.StatusForbidden, &api.Error{
				Code: -2,
				Message: "can't delete or edit this post",
			})
			return
		}

		ctx.Next()
	}
}