package feature

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func FavoriteActon(c *gin.Context) {
	//获取视频信息
	vID := c.Query("video_id")
	videoID, _ := strconv.Atoi(vID)
	//获取用户信息
	token := c.Query("token")
	var user User
	DB.Table("users").Where("token = ?", token).Find(&user)
	//更新视频点赞数
	DB.Exec("update videos set favorite_count = favorite_count+1 where id = ?", videoID)
	//更新点赞表
	DB.Table("favorites").Create(&Favorite{
		UserID:  user.ID,
		VideoID: videoID,
	})

	c.JSON(http.StatusOK, LikeAction{
		StatusCode: 0,
		StatusMsg:  "string",
	})
}

func FavoriteList(c *gin.Context) {
	//获取用户信息
	token := c.PostForm("token")
	var user User
	DB.Table("users").Where("token = ?", token).Find(&user)
	//获取点赞信息
	var favorites []Favorite
	DB.Table("favorites").Where("user_id = ?", user.ID).Find(&favorites)
	//获取视频信息
	videoIDs := make([]int, len(favorites))
	for _, f := range favorites {
		videoIDs = append(videoIDs, f.VideoID)
	}
	var videos []Video
	DB.Table("videos").Where("id in (?)", videoIDs).Find(&videos)

	c.JSON(http.StatusOK, Video_Feed{
		StatusCode: 0,
		StatusMsg:  "string",
		NextTime:   time.Now().Unix(),
		Vlist:      videos,
	})
}
