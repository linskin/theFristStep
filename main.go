package main

import (
	"fmt"

	"example.com/m/v2/feature"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	apiRouter := r.Group("/douyin")
	//基础接口
	apiRouter.GET("/feed/", feature.Feed)                  //视频流接口
	apiRouter.GET("/publish/list/", feature.PublishList)   //发布列表
	apiRouter.POST("/user/register/", feature.UserRegiset) //用户注册
	apiRouter.POST("/user/login/", feature.UserLogin)      //用户登录
	apiRouter.GET("/user/", feature.ShowUser)              //用户信息
	apiRouter.POST("/publish/action/", feature.PublishAction)
	//互动接口
	apiRouter.POST("/favorite/action/", feature.FavoriteActon) //赞操作
	apiRouter.GET("/comment/list/", feature.ShowComment)       //评论列表
	apiRouter.POST("/comment/action/", feature.CommentAction)
	//喜欢列表
	apiRouter.GET("/favorite/list/", feature.FavoriteList) //喜欢列表
}

func main() {
	fmt.Println("hello")

	r := gin.Default()
	initRouter(r)

	r.Run("192.168.1.102:8080")
}
