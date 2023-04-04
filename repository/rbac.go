package repository

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/luozi-csu/lzblogs/model"
	"github.com/pkg/errors"
)

type rbacRepository struct {
	enforcer *casbin.Enforcer
}

func newRBACRepository(modelConf string, a *gormadapter.Adapter) (RBACRepository, error) {
	e, err := casbin.NewEnforcer(modelConf, a)
	if err != nil {
		return nil, err
	}
	return &rbacRepository{
		enforcer: e,
	}, nil
}

func (rbac *rbacRepository) GetPolicies() []model.Policy {
	rules := rbac.enforcer.GetPolicy()
	policies := make([]model.Policy, 0, len(rules))
	for _, rule := range rules {
		policy := model.Policy{
			Subject: rule[0],
			Object:  rule[1],
			Action:  model.Operation(rule[2]),
		}
		policies = append(policies, policy)
	}
	return policies
}

func (rbac *rbacRepository) GetRoles() []model.Role {
	rules := rbac.enforcer.GetGroupingPolicy()
	roles := make([]model.Role, 0, len(rules))
	for _, rule := range rules {
		role := model.Role{
			Subject: rule[0],
			Name:    rule[1],
		}
		roles = append(roles, role)
	}
	return roles
}

func (rbac *rbacRepository) AddPolicy(policy *model.Policy) (*model.Policy, error) {
	rule := convertPolicyToSlice(policy)

	ok, err := rbac.enforcer.AddPolicy(rule...)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("policy already exists")
	}

	return policy, nil
}

func (rbac *rbacRepository) AddRole(role *model.Role) (*model.Role, error) {
	rule := convertRoleToSlice(role)

	ok, err := rbac.enforcer.AddGroupingPolicy(rule...)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("role already exists")
	}

	return role, nil
}

func (rbac *rbacRepository) AddPolicies(policies []model.Policy) ([]model.Policy, error) {
	var rules [][]string

	for i := 0; i < len(policies); i++ {
		rules = append(rules, convertPolicyToStrSlice(&policies[i]))
	}
	ok, err := rbac.enforcer.AddPolicies(rules)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("some policies already exist")
	}

	return policies, nil
}

func (rbac *rbacRepository) AddRoles(roles []model.Role) ([]model.Role, error) {
	var rules [][]string

	for i := 0; i < len(roles); i++ {
		rules = append(rules, convertRoleToStrSlice(&roles[i]))
	}
	ok, err := rbac.enforcer.AddGroupingPolicies(rules)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("some roles already exist")
	}

	return roles, nil
}

func (rbac *rbacRepository) UpdatePolicy(old, new *model.Policy) (*model.Policy, error) {
	oldRule, newRule := convertPolicyToStrSlice(old), convertPolicyToStrSlice(new)
	ok, err := rbac.enforcer.UpdatePolicy(oldRule, newRule)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("old policy not found")
	}

	return new, nil
}

func (rbac *rbacRepository) UpdateRole(old, new *model.Role) (*model.Role, error) {
	oldRule, newRule := convertRoleToStrSlice(old), convertRoleToStrSlice(new)
	ok, err := rbac.enforcer.UpdateGroupingPolicy(oldRule, newRule)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("old role not found")
	}

	return new, nil
}

func (rbac *rbacRepository) RemovePolicy(policy *model.Policy) error {
	rule := convertPolicyToSlice(policy)
	ok, err := rbac.enforcer.RemovePolicy(rule...)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("policy not found")
	}

	return nil
}

func (rbac *rbacRepository) RemoveRole(role *model.Role) error {
	rule := convertRoleToSlice(role)
	ok, err := rbac.enforcer.RemoveGroupingPolicy(rule...)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("role not found")
	}

	return nil
}

func (rbac *rbacRepository) HasPermission(user *model.User, permission ...string) (bool, error) {
	if user == nil || len(permission) == 0 {
		return false, errors.New("empty user or permission")
	}

	if user.Name == "" {
		return false, errors.New("empty username")
	}

	return rbac.enforcer.HasPermissionForUser(user.Name, permission...), nil
}

func convertPolicyToSlice(policy *model.Policy) []interface{} {
	if policy == nil {
		return nil
	}
	return []interface{}{policy.Subject, policy.Object, string(policy.Action)}
}

func convertRoleToSlice(role *model.Role) []interface{} {
	if role == nil {
		return nil
	}
	return []interface{}{role.Subject, role.Name}
}

func convertPolicyToStrSlice(policy *model.Policy) []string {
	if policy == nil {
		return nil
	}
	return []string{policy.Subject, policy.Object, string(policy.Action)}
}

func convertRoleToStrSlice(role *model.Role) []string {
	if role == nil {
		return nil
	}
	return []string{role.Subject, role.Name}
}
