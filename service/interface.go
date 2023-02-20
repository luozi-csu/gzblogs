package service

import (
	"github.com/luozi-csu/lzblogs/model"
)

type UserService interface {
	Get(string) (*model.User, error)
	List() (model.Users, error)
	Create(*model.CreateUserInput) (*model.User, error)
	Update(string, *model.UpdateUserInput) (*model.User, error)
	Delete(string) error
}

type RBACService interface {
	GetPolicies() []model.Policy
	GetRoles() []model.Role
	CreatePolicy(*model.Policy) (*model.Policy, error)
	CreateRole(*model.Role) (*model.Role, error)
	CreatePolicies([]model.Policy) ([]model.Policy, error)
	CreateRoles([]model.Role) ([]model.Role, error)
	UpdatePolicy(old, new *model.Policy) (*model.Policy, error)
	UpdateRole(old, new *model.Role) (*model.Role, error)
	DeletePolicy(*model.Policy) error
	DeleteRole(*model.Role) error
}
