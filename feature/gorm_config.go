package feature

import (
	"github.com/jinzhu/gorm"
)

// DB 创建gorm.DB对象的时候连接并没有被创建，在具体使用的时候才会创建。
// gorm内部(底层通过sql.DB实现)，准确的说是database/sql内部会维护一个连接池，可以通过参数设置最大空闲连接数，连接最大空闲时间等。
// 使用者不需要管连接的创建和关闭。
var DB *gorm.DB

func InitDB() {
	dsn := "remote:guoyixing@(119.23.244.10)/douyin?charset=utf8mb4&parseTime=True&Local"
	var err error
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	sqlDB := DB.DB()
	sqlDB.SetMaxIdleConns(10)  //总是存活连接数
	sqlDB.SetMaxOpenConns(200) //最大连接数
}
