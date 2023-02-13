package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"autoIncrement;primaryKey"`
	Name      string    `json:"name" gorm:"size:100;not null;unique"`
	Password  string    `json:"-" gorm:"size:256"`
	Avatar    string    `json:"avatar" gorm:"size:256"`
	Sign      string    `json:"sign" gorm:"size:256"`
	Phone     string    `json:"phone" gorm:"size:256"`
	Email     string    `json:"email" gorm:"size:256"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Roles     []Role    `json:"roles" gorm:"many2many:user_roles"`
}

type Users []User
