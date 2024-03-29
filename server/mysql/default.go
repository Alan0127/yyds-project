package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"yyds-pro/model"
)

var Client *gorm.DB

//初始化mysql连接
func InitMysql(config model.AppConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.App.Database.User,
		config.App.Database.Password,
		config.App.Database.Address,
		config.App.Database.DbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: New(),
	})

	if err != nil {
		return
	}
	Client = db
	return
}
