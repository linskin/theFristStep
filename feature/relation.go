package feature

import (
	"example.com/m/v2/conf"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func RelationAction(c *gin.Context) {
	Token := c.Query("token")
	conf.DB.AutoMigrate(&Follow{})
	var u User
	var tou User
	var count Follow
	var f Follow
	conf.DB.Where("token = ?", Token).First(&u)
	sid := c.Query("to_user_id")
	tuid, _ := strconv.Atoi(sid)
	conf.DB.Where("id = ?", tuid).First(&tou)
	at := c.Query("action_type")
	conf.DB.Model(&Follow{}).Last(&count)
	if at == "1" {
		u.FollowCount++
		tou.FollowerCount++
		//DB.Table("users").Select("follow_count").Where("id = ?", u.ID).Update(&u.FollowCount)
		//DB.Table("users").Select("follow_count").Where("id = ?", tou.ID).Update(&tou.FollowerCount)
		conf.DB.Table("users").Where("id = ?", u.ID).Update(User{FollowCount: u.FollowCount})
		conf.DB.Table("users").Where("id = ?", tou.ID).Update(User{FollowerCount: tou.FollowerCount, IsFollow: true})
		f = Follow{count.Id + 1, u.ID, tou.ID}
		fmt.Println(f)
		conf.DB.Model(&Follow{}).Create(&f)
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "success",
		})
		return
	}
	u.FollowCount--
	tou.FollowerCount--
	// DB.Table("users").Select("follow_count").Where("id = ?", u.ID).Update(&u.FollowCount)
	// DB.Table("users").Select("follow_count").Where("id = ?", tou.ID).Update(&tou.FollowerCount)
	conf.DB.Table("users").Where("id = ?", u.ID).Update(User{FollowCount: u.FollowCount})
	conf.DB.Table("users").Where("id = ?", tou.ID).Update(User{FollowerCount: tou.FollowerCount, IsFollow: false})
	conf.DB.Model(&Follow{}).Where("Uid = ? and TUid = ?", u.ID, tou.ID).First(&f)
	conf.DB.Model(&Follow{}).Delete(&f)
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "success",
	})
}

func FollowList(c *gin.Context) {
	s_id := c.Query("user_id")
	Uid, _ := strconv.Atoi(s_id)
	conf.DB.AutoMigrate(&Follow{})
	//var u User
	var Ulist []User
	var Tuid []Follow
	conf.DB.Table("follows").Where("Uid = ?", Uid).Find(&Tuid)
	for _, v := range Tuid {
		var u User
		conf.DB.Table("users").Where("id = ?", v.TUid).First(&u)
		fmt.Println(v)
		fmt.Println(u)
		Ulist = append(Ulist, u)
	}
	fmt.Println(Ulist)
	c.JSON(http.StatusOK, UserResponse{
		StatusCode: "0",
		StatusMsg:  "success",
		UserList:   Ulist,
	})

}

func FollowerList(c *gin.Context) {
	s_id := c.Query("user_id")
	tuid, _ := strconv.Atoi(s_id)
	conf.DB.AutoMigrate(&Follow{})
	var Ulist []User
	var Uid []Follow
	conf.DB.Table("follows").Where("TUid = ?", tuid).Find(&Uid)
	for _, v := range Uid {
		var u User
		conf.DB.Table("users").Where("id = ?", v.Uid).First(&u)
		Ulist = append(Ulist, u)
	}
	fmt.Println(Ulist)
	c.JSON(http.StatusOK, UserResponse{
		StatusCode: "0",
		StatusMsg:  "success",
		UserList:   Ulist,
	})
}

func FriendList(c *gin.Context) {
	s_id := c.Query("user_id")
	uid, _ := strconv.Atoi(s_id)
	var Ulist []User
	var Flist []User
	var Tuid []Follow
	conf.DB.Table("follows").Where("uid = ?", uid).Find(&Tuid)
	for _, v := range Tuid {
		var u User
		conf.DB.Table("users").Where("id = ?", v.TUid).First(&u)
		Ulist = append(Ulist, u)
	}
	var Uid []Follow
	conf.DB.Table("follows").Where("TUid = ?", uid).Find(&Uid)
	for _, v := range Uid {
		for _, uv := range Ulist {
			fmt.Println("1", uv)
			fmt.Println("2", v)
			if uv.ID == v.Uid {
				Flist = append(Flist, uv)
			}
		}
	}
	for _, F := range Flist {
		K1 := fmt.Sprintf("%d_%d", F.ID, uid)
		K2 := fmt.Sprintf("%d_%d", F.ID, uid)
		M_count.Store(K1, 0)
		M_count.Store(K2, 0)
	}
	c.JSON(http.StatusOK, UserResponse{
		StatusCode: "0",
		StatusMsg:  "success",
		UserList:   Flist,
	})
}
