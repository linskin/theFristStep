package feature

import (
	"example.com/m/v2/conf"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func ShowComment(c *gin.Context) {

	conf.DB.AutoMigrate(&User{}, &Video{}, &Comment{})
	vid := c.Query("video_id")
	num, _ := strconv.Atoi(vid)
	var comments []Comment
	conf.DB.Table("comments").Where("Vid = ?", num).Find(&comments)
	l := len(comments)
	conf.DB.Table("videos").Where("id = ?", num).Update("comment_count", l)
	for i, cm := range comments {
		var u_cm User
		conf.DB.Table("users").Where("id = ?", cm.Uid).Find(&u_cm)
		comments[i].User = u_cm
	}
	c.JSON(http.StatusOK, CommentResponse{
		StatusCode:  0,
		StatusMsg:   "success",
		CommentList: comments,
	})
}

func CommentAction(c *gin.Context) {
	conf.DB.AutoMigrate(&User{}, &Video{}, &Comment{})
	Token := c.Query("token")
	var com Comment
	var u User
	conf.DB.Table("users").Where("token = ?", Token).First(&u)
	Action_type := c.Query("action_type")
	vid := c.Query("video_id")
	num, _ := strconv.Atoi(vid)
	if Action_type == "1" {
		c_text := c.Query("comment_text")
		com = Comment{
			ID:         3,
			User:       u,
			Content:    c_text,
			CreateDate: time.RubyDate,
			Vid:        num,
			Uid:        u.ID,
		}
		conf.DB.Table("comments").Create(&com)
		//DB.Table("videos").Where("id = ?",num).Update("")
		c.JSON(http.StatusOK, Commentaction{
			StatusCode: 0,
			StatusMsg:  "success",
			Comment:    com,
		})
		return
	}
	c_id := c.Query("comment_id")
	cid, _ := strconv.Atoi(c_id)
	conf.DB.Table("comments").Where("id = ?", cid).First(&com)
	conf.DB.Table("comments").Delete(&com)
	c.JSON(http.StatusOK, Response{StatusCode: 0})
}
