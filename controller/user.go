package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/luozi-csu/lzblogs/service"
)

type UserController struct {
	userService service.UserService
}

func (u *UserController) RegisterRoute(api *gin.RouterGroup) {

}
