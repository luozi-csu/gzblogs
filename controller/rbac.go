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

// @Summary List policies
// @Description List rbac policies
// @Produce json
// @Tags rbac
// @Security JWT
// @Success 200 {object} common.Response{data=[]model.Policy}
// @Router /api/v1/rbac/policies [get]
func (rbac *RBACController) GetPolicies(c *gin.Context) {
	policies := rbac.rbacService.GetPolicies()
	common.ResponseSuccess(c, policies)
}

// @Summary List roles
// @Description List rbac roles
// @Produce json
// @Tags rbac
// @Security JWT
// @Success 200 {object} common.Response{data=[]model.Role}
// @Router /api/v1/rbac/roles [get]
func (rbac *RBACController) GetRoles(c *gin.Context) {
	roles := rbac.rbacService.GetRoles()
	common.ResponseSuccess(c, roles)
}

// @Summary Create policy
// @Description Create rbac policy
// @Accept json
// @Produce json
// @Tags rbac
// @Security JWT
// @Param policy body model.Policy true "policy info"
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

// @Summary Create role
// @Description Create rbac role
// @Accept json
// @Produce json
// @Tags rbac
// @Security JWT
// @Param role body model.Role true "role info"
// @Success 200 {object} common.Response{data=model.Role}
// @Router /api/v1/rbac/role [post]
func (rbac *RBACController) CreateRole(c *gin.Context) {
	role := new(model.Role)
	if err := c.BindJSON(role); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	role, err := rbac.rbacService.CreateRole(role)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	common.ResponseSuccess(c, role)
}

// @Summary Create policies
// @Description Create rbac policies
// @Accept json
// @Produce json
// @Tags rbac
// @Security JWT
// @Param policies body []model.Policy true "policies info"
// @Success 200 {object} common.Response{data=[]model.Policy}
// @Router /api/v1/rbac/policies [post]
func (rbac *RBACController) CreatePolicies(c *gin.Context) {
	inputs := new(model.CreatePoliciesInput)
	if err := c.BindJSON(inputs); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	policies, err := rbac.rbacService.CreatePolicies(inputs.Policies)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	common.ResponseSuccess(c, policies)
}

// @Summary Create roles
// @Description Create rbac roles
// @Accept json
// @Produce json
// @Tags rbac
// @Security JWT
// @Param roles body []model.Role true "roles info"
// @Success 200 {object} common.Response{data=[]model.Role}
// @Router /api/v1/rbac/roles [post]
func (rbac *RBACController) CreateRoles(c *gin.Context) {
	inputs := new(model.CreateRolesInput)
	if err := c.BindJSON(inputs); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	roles, err := rbac.rbacService.CreateRoles(inputs.Roles)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	common.ResponseSuccess(c, roles)
}

// @Summary Update policy
// @Description Update rbac policy
// @Accept json
// @Produce json
// @Tags rbac
// @Security JWT
// @Param old body model.Policy true "old policy info"
// @Param new body model.Policy true "new policy info"
// @Success 200 {object} common.Response{data=model.Policy}
// @Router /api/v1/rbac/policy [put]
func (rbac *RBACController) UpdatePolicy(c *gin.Context) {
	input := new(model.UpdatePolicyInput)
	if err := c.BindJSON(input); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	policy, err := rbac.rbacService.UpdatePolicy(input.Old, input.New)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	common.ResponseSuccess(c, policy)
}

// @Summary Update role
// @Description Update rbac role
// @Accept json
// @Produce json
// @Tags rbac
// @Security JWT
// @Param old body model.Role true "old role info"
// @Param new body model.Role true "new role info"
// @Success 200 {object} common.Response{data=model.Role}
// @Router /api/v1/rbac/role [put]
func (rbac *RBACController) UpdateRole(c *gin.Context) {
	input := new(model.UpdateRoleInput)
	if err := c.BindJSON(input); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	role, err := rbac.rbacService.UpdateRole(input.Old, input.New)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	common.ResponseSuccess(c, role)
}

// @Summary Delete policy
// @Description Delete rbac policy
// @Accept json
// @Tags rbac
// @Security JWT
// @Param policy body model.policy true "policy info"
// @Success 200 {object} common.Response{data=nil}
// @Router /api/v1/rbac/policy [delete]
func (rbac *RBACController) DeletePolicy(c *gin.Context) {
	policy := new(model.Policy)
	if err := c.BindJSON(policy); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	if err := rbac.rbacService.DeletePolicy(policy); err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	common.ResponseSuccess(c, nil)
}

// @Summary Delete role
// @Description Delete rbac role
// @Accept json
// @Tags rbac
// @Security JWT
// @Param role body model.Role true "role info"
// @Success 200 {object} common.Response{data=nil}
// @Router /api/v1/rbac/role [delete]
func (rbac *RBACController) DeleteRole(c *gin.Context) {
	role := new(model.Role)
	if err := c.BindJSON(role); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	if err := rbac.rbacService.DeleteRole(role); err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	common.ResponseSuccess(c, nil)
}

func (rbac *RBACController) RegisterRoute(rg *gin.RouterGroup) {
	rg.GET("/rbac/policies", rbac.GetPolicies)
	rg.GET("/rbac/roles", rbac.GetRoles)
	rg.POST("/rbac/policy", rbac.CreatePolicy)
	rg.POST("/rbac/role", rbac.CreateRole)
	rg.POST("/rbac/policies", rbac.CreatePolicies)
	rg.POST("/rbac/roles", rbac.CreateRoles)
	rg.PUT("/rbac/policy", rbac.UpdatePolicy)
	rg.PUT("/rbac/role", rbac.UpdateRole)
	rg.DELETE("/rbac/policy", rbac.DeletePolicy)
	rg.DELETE("/rbac/role", rbac.DeleteRole)
}
