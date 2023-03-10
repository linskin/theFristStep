package feature

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

/*
StatusCode int    `json:"status_code"`
StatusMsg  string `json:"status_msg"`
UserID     int    `json:"user_id"`
Token      string `json:"token"`
*/
func UserRegiset(c *gin.Context) {
	DB.AutoMigrate(&User{})
	var u User
	var count int
	name := c.Query("username")
	Password := c.Query("password")
	Token := "N" + name + "P" + Password
	result := DB.Where("Token = ?", Token).Find(&u)
	if result.Error == nil {
		c.JSON(http.StatusOK, UserLR{
			StatusCode: 1,
			StatusMsg:  "用户已经存在！请登录",
		})
	} else {
		DB.Table("users").Count(&count)
		DB.Create(&User{
			ID:       count + 1,
			Name:     name,
			IsFollow: false,
			Token:    Token,
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
	DB.AutoMigrate(&User{})
	name := "N" + c.Query("username")
	Password := "P" + c.Query("password")
	Token := name + Password
	var u User
	result := DB.Where("Token = ?", Token).Find(&u)
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
	DB.AutoMigrate(&User{})
	var u User
	DB.Where("Token = ?", Token).First(&u)
	c.JSON(http.StatusOK, UserInfo{
		StatusCode: 0,
		StatusMsg:  "个人信息",
		User:       u,
	})
}
