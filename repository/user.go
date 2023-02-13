package repository

import (
	"github.com/luozi-csu/lzblogs/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) GetUserByID(uint) (*model.User, error) {
	return nil, nil
}

func (u *userRepository) GetUserByAuthID(authType, authID string) (*model.User, error) {
	return nil, nil
}

func (u *userRepository) GetUserByName(string) (*model.User, error) {
	return nil, nil
}

func (u *userRepository) List() (model.Users, error) {
	return nil, nil
}

func (u *userRepository) Create(*model.User) (*model.User, error) {
	return nil, nil
}

func (u *userRepository) Update(*model.User) (*model.User, error) {
	return nil, nil
}

func (u *userRepository) Delete(*model.User) error {
	return nil
}

func (u *userRepository) AddRole(role *model.Role, user *model.User) error {
	return nil
}

func (u *userRepository) DelRole(role *model.Role, user *model.User) error {
	return nil
}

func (u *userRepository) Migrate() error {
	return u.db.AutoMigrate(&model.User{})
}
