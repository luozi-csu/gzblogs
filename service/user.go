package service

import (
	"strconv"

	"github.com/luozi-csu/lzblogs/model"
	"github.com/luozi-csu/lzblogs/repository"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) Get(id string) (*model.User, error) {
	return u.getUserByID(id)
}

func (u *userService) List() (model.Users, error) {
	return u.userRepository.List()
}

func (u *userService) Create(user *model.User) (*model.User, error) {
	// 使用bcrypt加密
	pwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "encryption failed")
	}
	user.Password = string(pwd)
	return u.userRepository.Create(user)
}

func (u *userService) Update(id string, new *model.User) (*model.User, error) {
	return nil, nil
}

func (u *userService) Delete(id string) error {
	return nil
}

func (u *userService) AddRole(id, rid string) error {
	return nil
}

func (u *userService) DelRole(id, rid string) error {
	return nil
}

func (u *userService) getUserByID(id string) (*model.User, error) {
	uid, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.Wrapf(err, "convert uid=%s from string to int failed", id)
	}
	return u.userRepository.GetUserByID(uint(uid))
}
