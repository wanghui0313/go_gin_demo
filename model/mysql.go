package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const DSN string = "root:root@tcp(127.0.0.1:3306)/vueshop?charset=utf8mb4"

var Db *gorm.DB
var dbErr error

func init() {
	Db, dbErr = gorm.Open(mysql.Open(DSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "shop_",
		},
		//DryRun: true,
	})
	if dbErr != nil {
		panic("failed to connect database")
	}
}
