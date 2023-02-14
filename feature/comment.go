package feature

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func ShowComment(c *gin.Context) {

	DB.AutoMigrate(&User{}, &Video{}, &Comment{})
	vid := c.Query("video_id")
	num, _ := strconv.Atoi(vid)
	var comments []Comment
	DB.Table("comments").Where("Vid = ?", num).Find(&comments)
	l := len(comments)
	DB.Table("videos").Where("id = ?", num).Update("comment_count", l)
	for i, cm := range comments {
		var u_cm User
		DB.Table("users").Where("id = ?", cm.Uid).Find(&u_cm)
		comments[i].User = u_cm
	}
	c.JSON(http.StatusOK, CommentResponse{
		StatusCode:  0,
		StatusMsg:   "success",
		CommentList: comments,
	})
}

func CommentAction(c *gin.Context) {
	DB.AutoMigrate(&User{}, &Video{}, &Comment{})

	//获取请求参数
	token := c.Query("token")            //token
	actionType := c.Query("action_type") //1评论 2删除评论
	vid := c.Query("video_id")
	num, _ := strconv.Atoi(vid) //视频id

	var com Comment
	var u User
	DB.Table("users").Where("token = ?", token).First(&u)

	if actionType == "1" {
		//新增评论
		cText := c.Query("comment_text")
		com = Comment{
			User:       u,
			Content:    cText,
			CreateDate: time.Now().Format("2006-01-02 15:04:05"),
			Vid:        num,
			Uid:        u.ID,
		}
		//插入数据库
		DB.Table("comments").Create(&com)
		//更新视频评论数
		DB.Exec("update videos set comment_count = comment_count+1 where id = ?", num)
		c.JSON(http.StatusOK, Commentaction{
			StatusCode: 0,
			StatusMsg:  "success",
			Comment:    com,
		})
		return
	}
	//删除评论
	cID := c.Query("comment_id")
	cid, _ := strconv.Atoi(cID)
	//删除评论记录
	DB.Table("comments").Delete(&Comment{}, cid)
	//减少视频评论数
	DB.Exec("update videos set comment_count = comment_count-1 where id = ?", num)
	c.JSON(http.StatusOK, Response{StatusCode: 0})
}
