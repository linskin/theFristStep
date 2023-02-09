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
	db.Table("videos").Where("id = ?", fk).Find(&PList)
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
	v := Video{Title: T,
		UID: u.V_key}
	db.Table("videos").Create(&v)
	filename := filepath.Base(V.Filename)
	finalName := fmt.Sprintf("%d_%s", u.ID, filename)
	saveFile := filepath.Join("./video_data", filename)
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
