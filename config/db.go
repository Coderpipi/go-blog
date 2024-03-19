package config

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Mysql struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

func initMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Cfg.Mysql.UserName, Cfg.Mysql.Password, Cfg.Mysql.Host,
		Cfg.Mysql.Port, Cfg.Mysql.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("init mysql error")
	}
	zap.L().Info("init mysql success")
	DB = db
}
