package feature

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func FavoriteActon(c *gin.Context) {
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
