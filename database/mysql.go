package database

import (
	"fmt"

	"github.com/luozi-csu/lzblogs/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
