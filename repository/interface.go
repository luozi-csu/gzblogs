package repository

import (
	"context"

	"github.com/luozi-csu/lzblogs/model"
)

type Repository interface {
	User() UserRepository
	RBAC() RBACRepository
	Ping(ctx context.Context) error
	Init() error
	Migrant
}

// migrant在server初始化时调用gorm自动迁移
type Migrant interface {
	Migrate() error
}

type UserRepository interface {
	GetUserByID(uint) (*model.User, error)
	GetUserByName(string) (*model.User, error)
	List() (model.Users, error)
	Create(*model.User) (*model.User, error)
	Update(*model.User) (*model.User, error)
	Delete(*model.User) error
	Migrate() error
}

type RBACRepository interface {
	GetPolicies() []model.Policy
	GetRoles() []model.Role
	AddPolicy(*model.Policy) (*model.Policy, error)
	AddRole(*model.Role) (*model.Role, error)
	AddPolicies([]model.Policy) ([]model.Policy, error)
	AddRoles([]model.Role) ([]model.Role, error)
	UpdatePolicy(old, new *model.Policy) (*model.Policy, error)
	UpdateRole(old, new *model.Role) (*model.Role, error)
	RemovePolicy(*model.Policy) error
	RemoveRole(*model.Role) error
}
