package repository

import (
	"context"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

type repository struct {
	db       *gorm.DB
	user     UserRepository
	rbac     RBACRepository
	migrants []Migrant
}

func NewRepository(db *gorm.DB, modelConf string, adapter *gormadapter.Adapter) Repository {
	r := &repository{
		db:   db,
		user: newUserRepository(db),
	}
	rbacRepository, _ := newRBACRepository(modelConf, adapter)
	r.rbac = rbacRepository

	r.migrants = getMigrants(r.user)
	return r
}

func getMigrants(objs ...interface{}) []Migrant {
	migrants := make([]Migrant, 0)
	for _, obj := range objs {
		if m, ok := obj.(Migrant); ok {
			migrants = append(migrants, m)
		}
	}
	return migrants
}

func (r *repository) Migrate() error {
	for _, m := range r.migrants {
		if err := m.Migrate(); err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) User() UserRepository {
	return r.user
}

func (r *repository) RBAC() RBACRepository {
	return r.rbac
}

// 数据库连接检查
func (r *repository) Ping(ctx context.Context) error {
	db, err := r.db.DB()
	if err != nil {
		return err
	}
	if err = db.PingContext(ctx); err != nil {
		return err
	}

	return nil
}

func (r *repository) Init() error {
	return nil
}
