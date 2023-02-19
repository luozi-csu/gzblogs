package service

import (
	"github.com/luozi-csu/lzblogs/model"
)

type UserService interface {
	Get(string) (*model.User, error)
	List() (model.Users, error)
	Create(*model.CreateUserInput) (*model.User, error)
	Update(string, *model.UpdateUserInput) (*model.User, error)
	Delete(string) error
}
