package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/goexam/internal/models"

	"github.com/goexam"
	"github.com/labstack/echo"
)

// UserController is return UserController
type UserController struct {
	goexam.UserService
}

// NewUserController is return UserController
func NewUserController(userSvc goexam.UserService) *UserController {
	return &UserController{
		userSvc,
	}
}

// Login is
func (uc *UserController) Login(ctx echo.Context) error {
	request := new(models.UserLoginRequest)
	err := ctx.Bind(request)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	err = uc.UserService.Login(request.Username, request.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "帐号或密码错误")
	}
	return ctx.JSON(http.StatusOK, "登录成功")
}

// Create is
func (uc *UserController) Create(ctx echo.Context) error {
	user := new(goexam.User)
	err := ctx.Bind(user)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	err = uc.UserService.Create(user)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "失败")
	}
	return ctx.NoContent(http.StatusOK)
}

// Delete is
func (uc *UserController) Delete(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	if id == 0 || err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	err = uc.UserService.Delete(uint(id))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "系统错误")
	}
	return ctx.NoContent(http.StatusOK)
}

// Update is
func (uc *UserController) Update(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	if id == 0 || err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	user := new(goexam.User)
	err = ctx.Bind(user)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	user.ID = uint(id)
	err = uc.UserService.Update(user)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "系统错误")
	}
	return ctx.NoContent(http.StatusOK)
}

// Get is
func (uc *UserController) Get(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	if id == 0 || err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	user, err := uc.UserService.Get(uint(id))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "系统错误")
	}
	userRes := new(models.UserResponse)
	log.Println(&userRes)
	userRes.BuildResponse(user)
	userRes.BuildResponse(user)
	log.Println(&userRes)
	return ctx.JSON(http.StatusOK, userRes)
}

// GetList is
func (uc *UserController) GetList(ctx echo.Context) error {
	filter := new(goexam.UserFilter)
	err := ctx.Bind(filter)
	filter.LoadDefault()
	if err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	userList, err := uc.UserService.GetList(filter)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "系统错误")
	}
	resUsers := make([]models.UserResponse, len(userList))
	for i := 0; i < len(userList); i++ {
		resUsers[i].BuildResponse(userList[i])
	}
	return ctx.JSON(http.StatusOK, resUsers)
}
