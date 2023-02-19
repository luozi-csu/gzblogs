package common

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luozi-csu/lzblogs/utils/logx"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(code, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	NewResponse(c, http.StatusOK, data, "success")
}

func ResponseFailed(c *gin.Context, code int, err error) {
	if code == 0 {
		code = http.StatusInternalServerError
	}
	if code == http.StatusUnauthorized && c.Request != nil {
		if val, err := c.Cookie(CookieTokenName); err == nil && val != "" {
			c.SetCookie(CookieTokenName, "", -1, "/", "", true, true)
			c.SetCookie(CookieLoginUser, "", -1, "/", "", true, false)
		}
	}

	var msg, userName string
	user := GetUser(c)
	if user != nil {
		userName = user.Name
	}

	if userName == "" {
		msg = fmt.Sprintf("user=unknown request=[%s %s] error=%v", c.Request.Method, c.Request.URL.Path, err)
	} else {
		msg = fmt.Sprintf("user=%s request=[%s %s] error=%v", userName, c.Request.Method, c.Request.URL.Path, err)
	}
	logx.Errorf(msg)

	NewResponse(c, code, nil, msg)
}
