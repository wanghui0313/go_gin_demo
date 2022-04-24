package admin

import (
	"server/model"
	_ "server/model"
)

var db = model.Db

type Admin struct {
	Id         uint   `gorm:"type:int;primary key;autoIncrement;UNSIGNED"`
	Username   string `gorm:"type:varchar(20);not null;index;comment:用户名"`
	Password   string `gorm:"type:char(40);not null"`
	Email      string `gorm:"type:varchar(255)"`
	Age        uint   `gorm:"type:tinyint(4)"`
	CreateTime int64  `gorm:"autoCreateTime"`
	UpdateTime int64  `gorm:"autoUpdateTime"`
}

//func (a *Admin) TableName() string {
//	return "xiaoniu_admin"
//}

func AdminInit() {
	//db.AutoMigrate(&Admin{})
	//admin := Admin{Username: "admin", Password: "123456", Email: "1312104251@qq.com", Age: 20}
	//result := db.Select("Username", "Password").Create(&admin)
}

func GetAdmin() Admin {
	var admin Admin
	db.First(&admin)
	return admin
}
