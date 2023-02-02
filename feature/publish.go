package feature

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, Video_Feed{
		StatusCode: 0,
		StatusMsg:  "string",
		NextTime:   time.Now().Unix(),
		Vlist:      DemoVideo,
	})
}
