package repository

import (
	"github.com/luozi-csu/lzblogs/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var (
	userCreateField = []string{"name", "password"}
)

type userRepository struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) GetUserByID(id uint) (*model.User, error) {
	user := new(model.User)
	if err := u.db.Omit("Password").Preload("Roles").First(user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) GetUserByName(name string) (*model.User, error) {
	user := new(model.User)
	if err := u.db.Omit("Password").Preload("Roles").Where("name = ?", name).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) List() (model.Users, error) {
	users := make(model.Users, 0)
	if err := u.db.Omit("Password").Preload("Roles").Order("name").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) Create(user *model.User) (*model.User, error) {
	if err := u.db.Select(userCreateField).Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) Update(user *model.User) (*model.User, error) {
	if user == nil {
		return nil, errors.New("empty user")
	}

	if err := u.db.Model(&model.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) Delete(user *model.User) error {
	if err := u.db.Delete(user, user.ID).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepository) Migrate() error {
	return u.db.AutoMigrate(&model.User{})
}
