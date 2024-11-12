package sOrm

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shopkone-service/hack"
)

var db *gorm.DB

func NewDb(shopID *uint) *gorm.DB {
	db.Statement.Settings.Store("__SHOP_ID", shopID)
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

	d.Callback().Query().Before("gorm:query").Register("gorm:add_shop_id", func(d *gorm.DB) {
		id, ok := getShopID(d)
		if !ok {
			return
		}
		d.Where("shop_id = ?", id)
	})
	d.Callback().Create().Before("gorm:before_create").Register("gorm:create_shop_id", func(d *gorm.DB) {
		id, ok := getShopID(d)
		if !ok {
			return
		}
		d.Statement.SetColumn("shop_id", id)
	})
	d.Callback().Update().Before("gorm:before_update").Register("gorm:update_shop_id", func(d *gorm.DB) {
		id, ok := getShopID(d)
		if !ok {
			return
		}
		d.Where("shop_id = ?", id)
		d.Statement.SetColumn("shop_id", id)
	})
	d.Callback().Delete().Before("gorm:before_delete").Register("gorm:delete_shop_id", func(d *gorm.DB) {
		id, ok := getShopID(d)
		if !ok {
			return
		}
		d.Where("shop_id = ?", id)
		d.Statement.SetColumn("shop_id", id)
	})

	db = d.Debug()

	if err = db.AutoMigrate(migrate...); err != nil {
		return err
	}

	return nil
}

func getShopID(d *gorm.DB) (id interface{}, c bool) {
	id, ok := d.Get("__SHOP_ID")
	tableName := d.Statement.Table
	if tableName == "shops" {
		return id, false
	}
	if tableName == "users" {
		return id, false
	}
	if tableName == "staffs" {
		return id, false
	}
	if tableName == "user_columns" {
		return id, false
	}
	if !ok {
		panic("sorm.connect.getShopID 异常")
		return id, false
	}
	g.Dump(tableName)
	return id, true
}
