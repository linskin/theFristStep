package feature

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func ShowComment(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@(localhost)/douyin?charset=utf8mb4&parseTime=True&Local")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{}, &Video{}, &Comment{})
	vid := c.Query("video_id")
	num, _ := strconv.Atoi(vid)
	var comments []Comment
	db.Table("comments").Where("Vid = ?", num).Find(&comments)
	l := len(comments)
	db.Table("videos").Where("id = ?", num).Update("comment_count", l)
	for i, cm := range comments {
		var u_cm User
		db.Table("users").Where("id = ?", cm.Uid).Find(&u_cm)
		comments[i].User = u_cm
	}
	c.JSON(http.StatusOK, CommentResponse{
		StatusCode:  0,
		StatusMsg:   "success",
		CommentList: comments,
	})
}

func CommentAction(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@(localhost)/douyin?charset=utf8mb4&parseTime=True&Local")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{}, &Video{}, &Comment{})
	Token := c.Query("token")
	var com Comment
	var u User
	db.Table("users").Where("token = ?", Token).First(&u)
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
		db.Table("comments").Create(&com)
		//db.Table("videos").Where("id = ?",num).Update("")
		c.JSON(http.StatusOK, Commentaction{
			StatusCode: 0,
			StatusMsg:  "success",
			Comment:    com,
		})
		return
	}
	c_id := c.Query("comment_id")
	cid, _ := strconv.Atoi(c_id)
	db.Table("comments").Where("id = ?", cid).First(&com)
	db.Table("comments").Delete(&com)
	c.JSON(http.StatusOK, Response{StatusCode: 0})
}
