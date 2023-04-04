package model

// casbin规则模型
type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:64;uniqueIndex:unique_index"`
	V0    string `gorm:"size:64;uniqueIndex:unique_index"`
	V1    string `gorm:"size:64;uniqueIndex:unique_index"`
	V2    string `gorm:"size:64;uniqueIndex:unique_index"`
	V3    string `gorm:"size:64;uniqueIndex:unique_index"`
	V4    string `gorm:"size:64;uniqueIndex:unique_index"`
	V5    string `gorm:"size:64;uniqueIndex:unique_index"`
}

type Operation string

var (
	AllOperation    Operation = "*"
	GetOperation    Operation = "get"
	CreateOperation Operation = "post"
	UpdateOperation Operation = "put"
	RemoveOperation Operation = "delete"
	ListOperation   Operation = "list"
)

// casbin策略
type Policy struct {
	Subject string    `json:"subject"`
	Object  string    `json:"object"`
	Action  Operation `json:"action"`
}

// casbin角色
type Role struct {
	Subject string `json:"subject"`
	Name    string `json:"name"`
}

type UpdatePolicyInput struct {
	Old *Policy `json:"old"`
	New *Policy `json:"new"`
}

type UpdateRoleInput struct {
	Old *Role `json:"old"`
	New *Role `json:"new"`
}

type CreatePoliciesInput struct {
	Policies []Policy `json:"policies"`
}

type CreateRolesInput struct {
	Roles []Role `json:"roles"`
}

type Resource struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}
