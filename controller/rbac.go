package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luozi-csu/lzblogs/common"
	"github.com/luozi-csu/lzblogs/model"
	"github.com/luozi-csu/lzblogs/service"
)

type RBACController struct {
	rbacService service.RBACService
}

func NewRBACController(rbacService service.RBACService) Controller {
	return &RBACController{
		rbacService: rbacService,
	}
}

// @Summary Create policy
// @Description Create rbac policy
// @Accept json
// @Produce json
// @Tags user
// @Security JWT
// @Param user body model.Policy true "policy info"
// @Success 200 {object} common.Response{data=model.Policy}
// @Router /api/v1/rbac/policy [post]
func (rbac *RBACController) CreatePolicy(c *gin.Context) {
	policy := new(model.Policy)
	if err := c.BindJSON(policy); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	policy, err := rbac.rbacService.CreatePolicy(policy)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	common.ResponseSuccess(c, policy)
}

func (rbac *RBACController) RegisterRoute(rg *gin.RouterGroup) {
	rg.POST("/rbac/policy", rbac.CreatePolicy)
}
