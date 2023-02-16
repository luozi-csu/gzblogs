package repository

import (
	"context"

	"gorm.io/gorm"
)

type repository struct {
	db   *gorm.DB
	user UserRepository
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
		user: newUserRepository(db),
	}
}

func (r *repository) User() UserRepository {
	return r.user
}

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

func (r *repository) Migrate() error {
	return nil
}