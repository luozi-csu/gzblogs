package service

import (
	"github.com/luozi-csu/lzblogs/model"
)

type UserService interface {
	Get(string) (*model.User, error)
	List() (model.Users, error)
	Create(*model.User) (*model.User, error)
	Update(string, *model.User) (*model.User, error)
	Delete(string) error
}
