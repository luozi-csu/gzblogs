package database

import (
	"fmt"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/luozi-csu/lzblogs/config"
	"github.com/luozi-csu/lzblogs/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func NewMysqlAdapter(db *gorm.DB) (*gormadapter.Adapter, error) {
	return gormadapter.NewAdapterByDBWithCustomTable(db, &model.CasbinRule{})
}
