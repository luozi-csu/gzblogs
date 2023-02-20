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
	AllOperation   Operation = "*"
	ReadOperation  Operation = "read"
	WriteOperation Operation = "write"
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

type CreatePolicyInput struct {
	Subject string    `json:"subject"`
	Object  string    `json:"object"`
	Action  Operation `json:"action"`
}

func (p *CreatePolicyInput) GetPolicy() *Policy {
	return &Policy{
		Subject: p.Subject,
		Object:  p.Object,
		Action:  p.Action,
	}
}

type CreateRoleInput struct {
	Subject string `json:"subject"`
	Name    string `json:"name"`
}

func (r *CreateRoleInput) GetRole() *Role {
	return &Role{
		Subject: r.Subject,
		Name:    r.Name,
	}
}
