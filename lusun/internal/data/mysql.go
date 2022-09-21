package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlRepo struct {
	Db *gorm.DB
}

// NewMysqlRepo .
func NewMysqlRepo() *MysqlRepo {
	dsn := "root:lu2252897@tcp(175.24.152.229:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}

	return &MysqlRepo{Db: db}
}
