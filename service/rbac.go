package service

import (
	"github.com/luozi-csu/lzblogs/model"
	"github.com/luozi-csu/lzblogs/repository"
)

type rbacService struct {
	rbacRepository repository.RBACRepository
}

func NewRBACService(rbacRepository repository.RBACRepository) RBACService {
	return &rbacService{
		rbacRepository: rbacRepository,
	}
}

func (rbac *rbacService) GetPolicies() []model.Policy {
	return rbac.rbacRepository.GetPolicies()
}

func (rbac *rbacService) GetRoles() []model.Role {
	return rbac.rbacRepository.GetRoles()
}

func (rbac *rbacService) CreatePolicy(policy *model.Policy) (*model.Policy, error) {
	return rbac.rbacRepository.AddPolicy(policy)
}

func (rbac *rbacService) CreateRole(role *model.Role) (*model.Role, error) {
	return rbac.rbacRepository.AddRole(role)
}

func (rbac *rbacService) CreatePolicies(policies []model.Policy) ([]model.Policy, error) {
	return rbac.rbacRepository.AddPolicies(policies)
}

func (rbac *rbacService) CreateRoles(roles []model.Role) ([]model.Role, error) {
	return rbac.rbacRepository.AddRoles(roles)
}

func (rbac *rbacService) UpdatePolicy(old, new *model.Policy) (*model.Policy, error) {
	return rbac.rbacRepository.UpdatePolicy(old, new)
}

func (rbac *rbacService) UpdateRole(old, new *model.Role) (*model.Role, error) {
	return rbac.rbacRepository.UpdateRole(old, new)
}

func (rbac *rbacService) DeletePolicy(policy *model.Policy) error {
	return rbac.rbacRepository.RemovePolicy(policy)
}

func (rbac *rbacService) DeleteRole(role *model.Role) error {
	return rbac.rbacRepository.RemoveRole(role)
}
