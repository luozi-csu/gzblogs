package model

import "time"

type User struct {
	ID        uint       `json:"id" gorm:"autoIncrement;primaryKey"`
	Name      string     `json:"name" gorm:"size:100;not null;unique"`
	Password  string     `json:"-" gorm:"size:256"`
	Avatar    string     `json:"avatar" gorm:"size:256"`
	Sign      string     `json:"sign" gorm:"size:256"`
	Email     string     `json:"email" gorm:"size:256"`
	AuthInfos []AuthInfo `json:"auth_infos" gorm:"foreignKey:UserID;references:ID"`
	BaseModel
}

type Users []User

// 重写gorm表名
func (*User) TableName() string {
	return "users"
}

type CreateUserInput struct {
	Name      string     `json:"name"`
	Password  string     `json:"password"`
	Email     string     `json:"email"`
	Avatar    string     `json:"avatar"`
	AuthInfos []AuthInfo `json:"auth_infos"`
}

func (u *CreateUserInput) GetUser() *User {
	return &User{
		Name:      u.Name,
		Password:  u.Password,
		Email:     u.Email,
		Avatar:    u.Avatar,
		AuthInfos: u.AuthInfos,
	}
}

type UpdateUserInput struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Sign     string `json:"sign"`
	Email    string `json:"email"`
}

func (u *UpdateUserInput) GetUser() *User {
	return &User{
		ID:       u.ID,
		Name:     u.Name,
		Password: u.Password,
		Avatar:   u.Avatar,
		Sign:     u.Sign,
		Email:    u.Email,
	}
}

type AuthInfo struct {
	ID           uint          `json:"id" gorm:"autoIncrement;primaryKey"`
	UserID       uint          `json:"uid"`
	AuthID       string        `json:"name" gorm:"size:256"`
	Url          string        `json:"url" gorm:"size:256"`
	AuthType     string        `json:"auth_type" gorm:"size:256"`
	AccessToken  string        `json:"-" gorm:"size:256"`
	RefreshToken string        `json:"-" gorm:"size:256"`
	Expiry       time.Duration `json:"-"`
	BaseModel
}

func (*AuthInfo) TableName() string {
	return "auth_infos"
}

type AuthUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	AuthType string `json:"auth_type"`
	AuthCode string `json:"auth_code"`
}
