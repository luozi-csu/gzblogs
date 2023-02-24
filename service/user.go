package service

import (
	"fmt"
	"strconv"

	"github.com/luozi-csu/lzblogs/model"
	"github.com/luozi-csu/lzblogs/repository"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

func (u *userService) Create(user *model.User) (*model.User, error) {
	if err := u.validateUser(user); err != nil {
		return nil, err
	}

	// 使用bcrypt加密
	pwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "encryption failed")
	}
	user.Password = string(pwd)
	return u.userRepository.Create(user)
}

func (u *userService) CreateOAuthUser(user *model.User) (*model.User, error) {
	if err := u.validateUser(user); err != nil {
		return nil, err
	}

	if len(user.AuthInfos) == 0 {
		return nil, errors.New("empty auth infos")
	}

	authInfo := user.AuthInfos[0]
	old, err := u.userRepository.GetUserByAuthID(authInfo.AuthType, authInfo.AuthID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.userRepository.Create(user)
		}
		return nil, err
	}

	return old, nil
}

func (u *userService) Auth(authUser *model.AuthUser) (*model.User, error) {
	if authUser == nil || authUser.Name == "" || authUser.Password == "" {
		return nil, errors.New("empty value in name, password")
	}

	user, err := u.userRepository.GetUserByName(authUser.Name)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authUser.Password)); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) Update(id string, new *model.User) (*model.User, error) {
	old, err := u.getUserByID(id)
	if err != nil {
		return nil, err
	}

	if new == nil {
		return nil, errors.New("emtpy user")
	}

	if new.ID != 0 && old.ID != new.ID {
		return nil, fmt.Errorf("update user=%s not match", id)
	}
	new.ID = old.ID

	if len(new.Password) > 0 {
		pwd, err := bcrypt.GenerateFromPassword([]byte(new.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, errors.Wrap(err, "encryption failed")
		}
		new.Password = string(pwd)
	}

	return u.userRepository.Update(new)
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

func (u *userService) validateUser(user *model.User) error {
	if user == nil {
		return errors.New("empty user")
	}
	if user.Name == "" {
		return errors.New("empty username")
	}
	if len(user.Password) < MinPasswordLength {
		return errors.New("password length is less than 6")
	}
	return nil
}
