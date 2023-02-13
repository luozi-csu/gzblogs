package controller

import "github.com/gin-gonic/gin"

type Contorller interface {
	RegisterRoute(*gin.RouterGroup)
}