package service

import (
	"github.com/luozi-csu/lzblogs/model"
	"github.com/luozi-csu/lzblogs/repository"
	"github.com/pkg/errors"
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
	if err := rbac.validatePolicy(policy); err != nil {
		return nil, err
	}
	return rbac.rbacRepository.AddPolicy(policy)
}

func (rbac *rbacService) CreateRole(role *model.Role) (*model.Role, error) {
	if err := rbac.validateRole(role); err != nil {
		return nil, err
	}
	return rbac.rbacRepository.AddRole(role)
}

func (rbac *rbacService) CreatePolicies(policies []model.Policy) ([]model.Policy, error) {
	n := len(policies)
	if n == 0 {
		return nil, errors.New("empty policies")
	}
	for i := 0; i < n; i++ {
		if err := rbac.validatePolicy(&policies[i]); err != nil {
			return nil, err
		}
	}
	return rbac.rbacRepository.AddPolicies(policies)
}

func (rbac *rbacService) CreateRoles(roles []model.Role) ([]model.Role, error) {
	n := len(roles)
	if n == 0 {
		return nil, errors.New("empty roles")
	}
	for i := 0; i < n; i++ {
		if err := rbac.validateRole(&roles[i]); err != nil {
			return nil, err
		}
	}
	return rbac.rbacRepository.AddRoles(roles)
}

func (rbac *rbacService) UpdatePolicy(old, new *model.Policy) (*model.Policy, error) {
	var err error
	if err = rbac.validatePolicy(old); err != nil {
		return nil, err
	}
	if err = rbac.validatePolicy(new); err != nil {
		return nil, err
	}
	return rbac.rbacRepository.UpdatePolicy(old, new)
}

func (rbac *rbacService) UpdateRole(old, new *model.Role) (*model.Role, error) {
	var err error
	if err = rbac.validateRole(old); err != nil {
		return nil, err
	}
	if err = rbac.validateRole(new); err != nil {
		return nil, err
	}
	return rbac.rbacRepository.UpdateRole(old, new)
}

func (rbac *rbacService) DeletePolicy(policy *model.Policy) error {
	return rbac.rbacRepository.RemovePolicy(policy)
}

func (rbac *rbacService) DeleteRole(role *model.Role) error {
	return rbac.rbacRepository.RemoveRole(role)
}

// 由于casbin创建和更新规则时不会忽略空值，因此必须对入参进行校验
func (rbac *rbacService) validatePolicy(policy *model.Policy) error {
	if policy == nil {
		return errors.New("empty policy")
	}
	if policy.Subject == "" || policy.Object == "" || policy.Action == "" {
		return errors.New("null value in subject, object, action")
	}
	return nil
}

func (rbac *rbacService) validateRole(role *model.Role) error {
	if role == nil {
		return errors.New("empty role")
	}
	if role.Subject == "" || role.Name == "" {
		return errors.New("null value in subject, name")
	}
	return nil
}
