package feature

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/*
StatusCode int    `json:"status_code"`
StatusMsg  string `json:"status_msg"`
UserID     int    `json:"user_id"`
Token      string `json:"token"`
*/
func UserRegiset(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@(localhost)/douyin?charset=utf8mb4&parseTime=True&Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{})
	var u User
	var count int
	name := c.Query("username")
	Password := c.Query("password")
	Token := "N" + name + "P" + Password
	result := db.Where("Token = ?", Token).Find(&u)
	if result.Error == nil {
		c.JSON(http.StatusOK, UserLR{
			StatusCode: 1,
			StatusMsg:  "用户已经存在！请登录",
		})
	} else {
		db.Table("users").Count(&count)
		db.Create(&User{
			ID:       count + 1,
			Name:     name,
			IsFollow: false,
			Token:    Token,
			V_key:    count + 1,
		})

		c.JSON(http.StatusOK, UserLR{
			StatusCode: 0,
			StatusMsg:  "注册成功！",
			UserID:     count + 1,
			Token:      Token,
		})
	}
}

func UserLogin(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@(localhost)/douyin?charset=utf8mb4&parseTime=True&Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{})
	name := "N" + c.Query("username")
	Password := "P" + c.Query("password")
	Token := name + Password
	var u User
	result := db.Where("Token = ?", Token).Find(&u)
	if result.Error == nil {
		c.JSON(http.StatusOK, UserLR{
			StatusCode: 0,
			StatusMsg:  "登录成功",
			UserID:     1,
			Token:      Token,
		})
	} else {
		c.JSON(http.StatusOK, UserLR{
			StatusCode: 1,
			StatusMsg:  "用户名或密码错误",
		})
	}
}

func ShowUser(c *gin.Context) {
	Token := c.Query("token")
	db, err := gorm.Open("mysql", "root:123456@(localhost)/douyin?charset=utf8mb4&parseTime=True&Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{})
	var u User
	db.Where("Token = ?", Token).Find(&u)
	c.JSON(http.StatusOK, UserInfo{
		StatusCode: 0,
		StatusMsg:  "个人信息",
		User:       u,
	})
}
