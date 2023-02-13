package feature

import (
	"example.com/m/v2/conf"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func FavoriteActon(c *gin.Context) {

	vid := c.Query("video_id")
	var v Video
	conf.DB.AutoMigrate(&User{}, &Video{})
	conf.DB.Table("videos").Where("id = ?", vid).First(&v)
	Likenum := v.FavoriteCount + 1
	conf.DB.Model(&v).Update(Video{FavoriteCount: Likenum})
	c.JSON(http.StatusOK, LikeAction{
		StatusCode: 0,
		StatusMsg:  "string",
	})
}

func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, Video_Feed{
		StatusCode: 0,
		StatusMsg:  "string",
		NextTime:   time.Now().Unix(),
		Vlist:      DemoVideo,
	})
}
