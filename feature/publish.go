package feature

import (
	"example.com/m/v2/conf"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func PublishList(c *gin.Context) {
	DB.AutoMigrate(&User{}, &Video{})
	//获取用户信息
	Token := c.Query("token")
	var u User
	DB.Table("users").Where("token = ?", Token).First(&u)
	//获取视频列表
	var PList []Video
	DB.Table("videos").Where("uid = ?", u.ID).Find(&PList)
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
	//获取用户信息
	var u User
	result := DB.Table("users").Where("Token = ?", Token).Find(&u)
	if result.Error != nil {
		panic(result.Error)
	}
	//获取表单文件
	V, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		panic(err)
	}
	//上传视频
	filename := filepath.Base(V.Filename)
	finalName := fmt.Sprintf("%d_%s_%s", u.ID, filename, time.Now().String())
	saveFile := filepath.Join("/douyin", filename)
	err = Upload(V, saveFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	}
	//储存视频对象
	v := Video{Title: T,
		PlayURL:  conf.CosUrl + saveFile,
		CoverURL: "https://img0.baidu.com/it/u=3294539948,324399065&fm=253&fmt=auto&app=138&f=JPEG?w=822&h=500",
		UID:      u.ID}
	DB.Table("videos").Create(&v)

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + "成功发布!",
	})
}
