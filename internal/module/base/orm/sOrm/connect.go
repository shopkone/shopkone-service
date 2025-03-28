package sOrm

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/driver/postgres"
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
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		conf.Mysql.Host,
		conf.Mysql.User,
		conf.Mysql.Password,
		conf.Mysql.Dbname,
		conf.Mysql.Port,
		conf.Mysql.TimeZone,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("初始化数据库失败: %w", err)
	}

	database.Callback().Query().Before("gorm:query").Register("gorm:add_shop_id", func(tx *gorm.DB) {
		id, ok := getShopID(tx)
		if !ok {
			return
		}
		tx = tx.Where("shop_id = ?", id)
	})
	database.Callback().Create().Before("gorm:before_create").Register("gorm:create_shop_id", func(tx *gorm.DB) {
		id, ok := getShopID(tx)
		if !ok {
			return
		}
		tx.Statement.SetColumn("shop_id", id)
	})
	database.Callback().Update().Before("gorm:before_update").Register("gorm:update_shop_id", func(tx *gorm.DB) {
		id, ok := getShopID(tx)
		if !ok {
			return
		}
		tx = tx.Where("shop_id = ?", id)
		tx.Statement.SetColumn("shop_id", id)
	})
	database.Callback().Delete().Before("gorm:before_delete").Register("gorm:delete_shop_id", func(tx *gorm.DB) {
		id, ok := getShopID(tx)
		if !ok {
			return
		}
		tx = tx.Where("shop_id = ?", id)
		tx.Statement.SetColumn("shop_id", id)
	})

	db = database.Debug()

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
	if tableName == "user_login_records" {
		return id, false
	}
	g.Dump(tableName, id)
	if !ok || id == 0 {
		panic("sorm.connect.getShopID 异常")
		return 0, true
	}
	return id, true
}
