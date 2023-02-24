package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luozi-csu/lzblogs/authentication"
	"github.com/luozi-csu/lzblogs/authorization/oauth"
	"github.com/luozi-csu/lzblogs/common"
	"github.com/luozi-csu/lzblogs/model"
	"github.com/luozi-csu/lzblogs/service"
)

type AuthController struct {
	userService  service.UserService
	jwtService   *authentication.JWTService
	oauthManager *oauth.OAuthManager
}

func NewAuthController(us service.UserService, js *authentication.JWTService, om *oauth.OAuthManager) *AuthController {
	return &AuthController{
		userService:  us,
		jwtService:   js,
		oauthManager: om,
	}
}

// @Summary Login
// @Description User login
// @Accept json
// @Produce json
// @Tags auth
// @Param user body model.AuthUser true "auth user info"
// @Success 200 {object} common.Response{data=model.JWTToken}
// @Router /api/v1/auth/login [post]
func (ac *AuthController) Login(c *gin.Context) {
	authUser := new(model.AuthUser)
	if err := c.BindJSON(authUser); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	var user *model.User
	var err error
	if authUser.AuthType != oauth.EmptyAuthType {
		client, err := ac.oauthManager.GetOAuthClient(authUser.AuthType)
		if err != nil {
			common.ResponseFailed(c, http.StatusBadRequest, err)
			return
		}

		token, err := client.GetToken(authUser.AuthCode)
		if err != nil {
			common.ResponseFailed(c, http.StatusBadRequest, err)
			return
		}

		userInfo, err := client.GetUserInfo(token)
		if err != nil {
			common.ResponseFailed(c, http.StatusBadRequest, err)
			return
		}

		user, err = ac.userService.CreateOAuthUser(userInfo.GetUser())
	} else {
		user, err = ac.userService.Auth(authUser)
	}
	if err != nil {
		common.ResponseFailed(c, http.StatusUnauthorized, err)
		return
	}

	token, err := ac.jwtService.NewToken(user)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	c.SetCookie(common.CookieTokenName, token, 24*3600, "/", "", true, true)
	c.SetCookie(common.CookieLoginUser, string(userJson), 24*3600, "/", "", true, true)

	common.ResponseSuccess(c, model.JWTToken{
		Token:       token,
		Description: "please set token in Header like [Authorization: Bearer <token>]",
	})
}

func (ac *AuthController) RegisterRoute(rg *gin.RouterGroup) {
	rg.POST("/auth/login", ac.Login)
}
