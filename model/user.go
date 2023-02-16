package model

type User struct {
	BaseModel
	ID       uint   `json:"id" gorm:"autoIncrement;primaryKey"`
	Name     string `json:"name" gorm:"size:100;not null;unique"`
	Password string `json:"-" gorm:"size:256"`
	Avatar   string `json:"avatar" gorm:"size:256"`
	Sign     string `json:"sign" gorm:"size:256"`
	Phone    string `json:"phone" gorm:"size:20"`
	Email    string `json:"email" gorm:"size:256"`
}

type Users []User

type CreatedUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

func (u *CreatedUser) GetUser() *User {
	return &User{
		Name:     u.Name,
		Password: u.Password,
		Email:    u.Email,
		Avatar:   u.Avatar,
	}
}
