package feature

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func FavoriteActon(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@(localhost)/douyin?charset=utf8mb4&parseTime=True&Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	vid := c.Query("video_id")
	var v Video
	db.AutoMigrate(&User{}, &Video{})
	db.Table("videos").Where("id = ?", vid).First(&v)
	Likenum := v.FavoriteCount+1
	db.Model(&v).Update(Video{FavoriteCount: Likenum})
	c.JSON(http.StatusOK, LikeAction{
		StatusCode: 0,
		StatusMsg:  "string",
	})
}
//
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, Video_Feed{
		StatusCode: 0,
		StatusMsg:  "string",
		NextTime:   time.Now().Unix(),
		Vlist:      DemoVideo,
	})
}
