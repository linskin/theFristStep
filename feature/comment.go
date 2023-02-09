package feature

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowComment(c *gin.Context) {
	c.JSON(http.StatusOK, CommentResponse{
		StatusCode:  0,
		StatusMsg:   "success",
		CommentList: DemoComment,
	})
}

func CommentAction(c *gin.Context) {
	Action_type := c.Query("action_type")
	if Action_type == "1" {
		c_text := c.Query("comment_text")
		com := Comment{
			ID:         3,
			User:       DemoAuthor,
			Content:    c_text,
			CreateDate: "2023-2-9",
		}
		c.JSON(http.StatusOK, Commentaction{
			StatusCode: 0,
			StatusMsg:  "success",
			Comment:    com,
		})
		return
	}
	c.JSON(http.StatusOK, Response{StatusCode: 0})
}
