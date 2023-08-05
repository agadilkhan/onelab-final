package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/agadilkhan/onelab-final/api"
	"github.com/agadilkhan/onelab-final/internal/entity"
	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags auth
// @Description create user
// @ID create user
// @Accept json
// @Produce json
// @Param input body api.RegisterRequest true "account info"
// @Succes 201 {object} api.Ok
// @Failure 400,404 {object} api.Error
// @Failure 500 {object} api.Error
// @Failure default {object} api.Error
// @Router /auth/register [post]
func (h *Handler) createUser(ctx *gin.Context) {
	var req api.RegisterRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
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
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

// @Summary SignIn
// @Tags auth
// @Description login user
// @ID login user
// @Accept json
// @Produce json
// @Param input body api.LoginRequest true "credentials"
// @Success 200 {string} api.Ok
// @Failure 400,404 {object} api.Error
// @Failure 500 {object} api.Error
// @Failure default {object} api.Error
// @Router /auth/login [post]
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
