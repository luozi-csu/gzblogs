package model

type Role struct {
	ID   string
	Name string
}

type Authorizer interface {
	HasPermission(userID, action, asset string) bool
}

type UserAuthorizer struct {
	users Users
	roles []Role
}

func (a *UserAuthorizer) HasPermission(userID, action, asset string) bool {
	return false
}
