package dao

import (
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

// "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"

func InitMySQL() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/bubble?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return DB.DB().Ping()
}
func Close() {
	DB.Close()
}
