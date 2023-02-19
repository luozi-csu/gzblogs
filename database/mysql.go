package database

import (
	"fmt"

	xormadapter "github.com/casbin/xorm-adapter/v2"
	"github.com/luozi-csu/lzblogs/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func NewMysqlAdapter(cfg *config.DatabaseConfig) (*xormadapter.Adapter, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/casbin?charset=utf8mb4&parseparseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port)
	return xormadapter.NewAdapter("mysql", dsn, true)
}
