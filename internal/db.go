package internal

import (
	"fmt"

	"github.com/yanun0323/pkg/logs"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB(conf Config) (*gorm.DB, error) {
	cfg := &gorm.Config{
		TranslateError: true,
	}

	if conf.SQLite {
		return gorm.Open(sqlite.Open("sqlite.db"), cfg)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.MySQL.Username,
		conf.MySQL.Password,
		conf.MySQL.Host,
		conf.MySQL.Port,
		conf.MySQL.Database)

	logs.Infof("mysql dsn: %s", dsn)

	return gorm.Open(mysql.Open(dsn), cfg)
}
