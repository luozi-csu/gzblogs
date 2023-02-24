package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luozi-csu/lzblogs/common"
	"github.com/luozi-csu/lzblogs/model"
	"github.com/luozi-csu/lzblogs/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) Controller {
	return &UserController{
		userService: userService,
	}
}

// @Summary Create user
// @Description Create user
// @Accept json
// @Produce json
// @Tags user
// @Security JWT
// @Param user body model.CreateUserInput true "user info"
// @Success 200 {object} common.Response{data=model.User}
// @Router /api/v1/users [post]
func (u *UserController) Create(c *gin.Context) {
	createUserInput := new(model.CreateUserInput)
	if err := c.BindJSON(createUserInput); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	user, err := u.userService.Create(createUserInput.GetUser())
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	common.ResponseSuccess(c, user)
}

// @Summary Get user
// @Description Get user
// @Produce json
// @Tags user
// @Security JWT
// @Param id path string true "user id"
// @Success 200 {object} common.Response{data=model.User}
// @Router /api/v1/users/{id} [get]
func (u *UserController) Get(c *gin.Context) {
	user, err := u.userService.Get(c.Param("id"))
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	common.ResponseSuccess(c, user)
}

// @Summary List users
// @Description List users
// @Produce json
// @Tags user
// @Security JWT
// @Success 200 {object} common.Response{data=[]model.User}
// @Router /api/v1/users [get]
func (u *UserController) List(c *gin.Context) {
	users, err := u.userService.List()
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	common.ResponseSuccess(c, users)
}

// @Summary Update user
// @Description Update user
// @Accept json
// @Produce json
// @Tags user
// @Security JWT
// @Param id path string true "user id"
// @Param user body model.UpdateUserInput true "new user info"
// @Success 200 {object} common.Response{data=model.User}
// @Router /api/v1/users/{id} [put]
func (u *UserController) Update(c *gin.Context) {
	updateUserInput := new(model.UpdateUserInput)
	if err := c.BindJSON(updateUserInput); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	user, err := u.userService.Update(c.Param("id"), updateUserInput.GetUser())
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	common.ResponseSuccess(c, user)
}

// @Summary Delete user
// @Description Delete user
// @Produce json
// @Tags user
// @Security JWT
// @Param id path string true "user id"
// @Success 200 {object} common.Response{data=nil}
// @Router /api/v1/users/{id} [delete]
func (u *UserController) Delete(c *gin.Context) {
	if err := u.userService.Delete(c.Param("id")); err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	common.ResponseSuccess(c, nil)
}

func (u *UserController) RegisterRoute(rg *gin.RouterGroup) {
	rg.POST("/users", u.Create)
	rg.GET("/users/:id", u.Get)
	rg.GET("/users", u.List)
	rg.PUT("/users/:id", u.Update)
	rg.DELETE("/users/:id", u.Delete)
}
