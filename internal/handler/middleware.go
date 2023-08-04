package handler

import (
	"errors"
	"net/http"
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