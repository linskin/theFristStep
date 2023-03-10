package feature

import (
	"example.com/m/v2/conf"
	"github.com/jinzhu/gorm"
)

// DB 创建gorm.DB对象的时候连接并没有被创建，在具体使用的时候才会创建。
// gorm内部(底层通过sql.DB实现)，准确的说是database/sql内部会维护一个连接池，可以通过参数设置最大空闲连接数，连接最大空闲时间等。
// 使用者不需要管连接的创建和关闭。
var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("mysql", conf.DSN)
	if err != nil {
		panic(err)
	}
	sqlDB := DB.DB()
	sqlDB.SetMaxIdleConns(10)  //总是存活连接数
	sqlDB.SetMaxOpenConns(200) //最大连接数
	DB.AutoMigrate(&User{}, &Favorite{}, &Video{}, &Comment{},
		&Follow{}, Message{})
}
