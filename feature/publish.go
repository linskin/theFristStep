package feature

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func PublishList(c *gin.Context) {
	Token := c.Query("token")
	db, err := gorm.Open("mysql", "root:123456@(localhost)/douyin?charset=utf8mb4&parseTime=True&Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{}, &Video{})
	var u User
	db.Table("users").Where("token = ?", Token).First(&u)
	fk := u.V_key
	var PList []Video
	db.Table("videos").Where("uid = ?", fk).Find(&PList)
	for _, v := range PList {
		v.Author = u
	}
	c.JSON(http.StatusOK, Video_Feed{
		StatusCode: 0,
		StatusMsg:  "string",
		NextTime:   time.Now().Unix(),
		Vlist:      PList,
	})
}

func PublishAction(c *gin.Context) {
	Token := c.PostForm("token")
	T := c.PostForm("title")
	db, err := gorm.Open("mysql", "root:123456@(localhost)/douyin?charset=utf8mb4&parseTime=True&Local")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{}, &Video{})
	var u User
	result := db.Table("users").Where("Token = ?", Token).Find(&u)
	if result.Error != nil {
		panic(err)
	}
	V, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		panic(err)
	}
	filename := filepath.Base(V.Filename)
	finalName := fmt.Sprintf("%d_%s", u.ID, filename)
	saveFile := filepath.Join("./public/", filename)
	v := Video{Title: T,
		PlayURL:  "https://192.168.1.102:8080/static/" + filename,
		CoverURL: "https://img0.baidu.com/it/u=3294539948,324399065&fm=253&fmt=auto&app=138&f=JPEG?w=822&h=500",
		UID:      u.V_key}
	db.Table("videos").Create(&v)
	if err := c.SaveUploadedFile(V, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  err.Error(),
		})
		panic(err)
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + "成功加载!",
	})
}
