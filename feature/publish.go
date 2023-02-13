package feature

import (
	"context"
	"example.com/m/v2/conf"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func PublishList(c *gin.Context) {
	Token := c.Query("token")
	conf.DB.AutoMigrate(&User{}, &Video{})
	var u User
	conf.DB.Table("users").Where("token = ?", Token).First(&u)
	fk := u.V_key
	var PList []Video
	conf.DB.Table("videos").Where("uid = ?", fk).Find(&PList)
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

	var u User
	result := conf.DB.Table("users").Where("Token = ?", Token).Find(&u)
	if result.Error != nil {
		panic(result.Error)
	}

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
		UID:      u.V_key}
	conf.DB.Table("videos").Create(&v)
	//if err := c.SaveUploadedFile(V, saveFile); err != nil {
	//	c.JSON(http.StatusOK, Response{
	//		StatusCode: 0,
	//		StatusMsg:  err.Error(),
	//	})
	//	panic(err)
	//}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + "成功发布!",
	})
}

func Upload(file *multipart.FileHeader, saveName string) error {
	fd, err := file.Open()
	if err != nil {
		return err
	}
	if _, err = conf.Client.Object.Put(context.Background(), saveName, fd, nil); err != nil {
		return err
	}
	return nil
}
