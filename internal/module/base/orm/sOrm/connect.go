package sOrm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shopkone-service/hack"
)

var db *gorm.DB

func NewDb() *gorm.DB {
	return db
}

func ConnectMysql(migrate []interface{}) error {
	conf, err := hack.GetConfig()
	if err != nil {
		return err
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		conf.Mysql.User,
		conf.Mysql.Password,
		conf.Mysql.Host,
		conf.Mysql.Port,
		conf.Mysql.Dbname,
		conf.Mysql.TimeZone,
	)

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("初始化数据库失败: %w", err)
	}

	db = d.Debug()

	if err = db.AutoMigrate(migrate...); err != nil {
		return err
	}

	return nil
}
