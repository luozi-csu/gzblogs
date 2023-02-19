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

// 重写gorm表名
func (*User) TableName() string {
	return "users"
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

func (u *CreateUserInput) GetUser() *User {
	return &User{
		Name:     u.Name,
		Password: u.Password,
		Email:    u.Email,
		Avatar:   u.Avatar,
	}
}

type UpdateUserInput struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Sign     string `json:"sign"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

func (u *UpdateUserInput) GetUser() *User {
	return &User{
		ID:       u.ID,
		Name:     u.Name,
		Password: u.Password,
		Avatar:   u.Avatar,
		Sign:     u.Sign,
		Phone:    u.Phone,
		Email:    u.Email,
	}
}
