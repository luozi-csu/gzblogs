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
// @Description Create user and storage
// @Accept json
// @Produce json
// @Tags user
// @Security JWT
// @Param user body model.CreatedUser true "user info"
// @Success 200 {object} common.Response{data=model.User}
// @Router /api/v1/users [post]
func (u *UserController) Create(c *gin.Context) {
	createdUser := new(model.CreatedUser)
	if err := c.BindJSON(createdUser); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	user := createdUser.GetUser()
	// if err := u.userService.Validate(user); err != nil {
	// 	common.ResponseFailed(c, http.StatusBadRequest, err)
	// 	return
	// }

	user, err := u.userService.Create(user)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	common.ResponseSuccess(c, user)
}

func (u *UserController) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/users", u.Create)
}
