package service

import (
	"fmt"
	"strconv"

	"github.com/luozi-csu/lzblogs/model"
	"github.com/luozi-csu/lzblogs/repository"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

const (
	MinPasswordLength = 6
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

func (u *userService) Create(input *model.CreateUserInput) (*model.User, error) {
	if input == nil {
		return nil, errors.New("empty user input")
	}
	if input.Name == "" {
		return nil, errors.New("empty user name")
	}
	if len(input.Password) < MinPasswordLength {
		return nil, fmt.Errorf("password length must great than %d", MinPasswordLength)
	}
	// 使用bcrypt加密
	pwd, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "encryption failed")
	}
	input.Password = string(pwd)
	return u.userRepository.Create(input.GetUser())
}

func (u *userService) Update(id string, new *model.UpdateUserInput) (*model.User, error) {
	old, err := u.getUserByID(id)
	if err != nil {
		return nil, err
	}

	if new == nil {
		return nil, errors.New("emtpy user input")
	}

	if new.ID != 0 && old.ID != new.ID {
		return nil, fmt.Errorf("update user=%s not match", id)
	}
	new.ID = old.ID

	return u.userRepository.Update(new.GetUser())
}

func (u *userService) Delete(id string) error {
	user, err := u.getUser(id)
	if err != nil {
		return err
	}

	return u.userRepository.Delete(user)
}

func (u *userService) getUserByID(id string) (*model.User, error) {
	uid, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.Wrapf(err, "convert uid=%s from string to int failed", id)
	}
	return u.userRepository.GetUserByID(uint(uid))
}

func (u *userService) getUser(id string) (*model.User, error) {
	uid, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.Wrapf(err, "convert uid=%s from string to int failed", id)
	}
	return &model.User{ID: uint(uid)}, nil
}
