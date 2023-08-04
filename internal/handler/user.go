package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/agadilkhan/onelab-final/api"
	"github.com/agadilkhan/onelab-final/internal/entity"
	"github.com/gin-gonic/gin"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (h *Handler) createUser(ctx *gin.Context) {
	var req api.RegisterRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	u := &entity.User{
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  req.Password,
	}

	err = h.service.Register(ctx, u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *Handler) loginUser(ctx *gin.Context) {
	var req api.LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	accessToken, err := h.service.Login(ctx, req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.Ok{
		Code:    0,
		Message: "success",
		Data:    accessToken,
	})
}

func getAuthUserID(ctx *gin.Context) (int, error) {

	userID, ok := ctx.MustGet(authUserID).(int)
	if !ok {
		log.Printf("can't get userID")
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "can't get user id from auth",
		})
		return -1, errors.New("can't get user id from auth")
	}

	return userID, nil
}
