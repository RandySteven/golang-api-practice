package handlers

import (
	"context"
	"net/http"
	"strconv"
	"test-api/entities/payload/req"
	"test-api/entities/payload/res"
	"test-api/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	usecase interfaces.UserUsecase
}

// GetUserById implements interfaces.UserHandler.
func (handler *UserHandler) GetUserById(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
		id        = c.Param("id")
	)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	userResp, err := handler.usecase.GetUserById(ctx, uint(idInt))
	if err != nil {
		return
	}
	resp := res.Response{
		Message: "Success get user",
		Data:    userResp,
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllUsers implements interfaces.UserHandler.
func (handler *UserHandler) GetAllUsers(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
	)
	usersRes, err := handler.usecase.GetAllUsers(ctx)
	if err != nil {
		return
	}
	res := res.Response{
		Message: "Success get all users",
		Data:    usersRes,
	}
	c.JSON(http.StatusOK, res)
}

// LoginUser implements interfaces.UserHandler.
func (handler *UserHandler) LoginUser(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
		login     *req.UserLogin
	)
	if err := c.ShouldBind(&login); err != nil {
		return
	}
	userRes, err := handler.usecase.LoginUser(ctx, login)
	if err != nil {
		return
	}
	resp := &res.Response{
		Message: "Success to login user",
		Data:    userRes,
	}
	c.JSON(http.StatusOK, resp)
}

// RegisterUser implements interfaces.UserHandler.
func (handler *UserHandler) RegisterUser(c *gin.Context) {
	var (
		requestId = uuid.NewString()
		ctx       = context.WithValue(c.Request.Context(), "request_id", requestId)
		register  *req.UserRegister
	)
	if err := c.ShouldBind(&register); err != nil {
		return
	}

	userRes, err := handler.usecase.RegisterUser(ctx, register)
	if err != nil {
		return
	}
	resp := res.Response{
		Message: "Success to register user",
		Data:    userRes,
	}
	c.JSON(http.StatusCreated, resp)
}

func NewUserHandler(usecase interfaces.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

var _ interfaces.UserHandler = &UserHandler{}
