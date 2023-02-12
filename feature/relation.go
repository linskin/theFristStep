package feature

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func RelationAction(c *gin.Context) {
	Token := c.Query("token")
	db, err := gorm.Open("mysql", "root:123456@(localhost)/douyin?charset=utf8mb4&parseTime=True&Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&Follow{})
	var u User
	var tou User
	var count Follow
	var f Follow
	db.Where("token = ?", Token).First(&u)
	sid := c.Query("to_user_id")
	tuid, _ := strconv.Atoi(sid)
	db.Where("id = ?", tuid).First(&tou)
	at := c.Query("action_type")
	db.Model(&Follow{}).Last(&count)
	if at == "1" {
		u.FollowCount++
		tou.FollowerCount++
		//db.Table("users").Select("follow_count").Where("id = ?", u.ID).Update(&u.FollowCount)
		//db.Table("users").Select("follow_count").Where("id = ?", tou.ID).Update(&tou.FollowerCount)
		db.Table("users").Where("id = ?", u.ID).Update(User{FollowCount: u.FollowCount})
		db.Table("users").Where("id = ?", tou.ID).Update(User{FollowerCount: tou.FollowerCount, IsFollow: true})
		f = Follow{count.Id + 1, u.ID, tou.ID}
		fmt.Println(f)
		db.Model(&Follow{}).Create(&f)
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "success",
		})
		return
	}
	u.FollowCount--
	tou.FollowerCount--
	// db.Table("users").Select("follow_count").Where("id = ?", u.ID).Update(&u.FollowCount)
	// db.Table("users").Select("follow_count").Where("id = ?", tou.ID).Update(&tou.FollowerCount)
	db.Table("users").Where("id = ?", u.ID).Update(User{FollowCount: u.FollowCount})
	db.Table("users").Where("id = ?", tou.ID).Update(User{FollowerCount: tou.FollowerCount, IsFollow: false})
	db.Model(&Follow{}).Where("Uid = ? and TUid = ?", u.ID, tou.ID).First(&f)
	db.Model(&Follow{}).Delete(&f)
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "success",
	})
}

func FollowList(c *gin.Context) {
	s_id := c.Query("user_id")
	Uid, _ := strconv.Atoi(s_id)
	db, err := gorm.Open("mysql", "root:123456@(localhost)/douyin?charset=utf8mb4&parseTime=True&Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&Follow{})
	//var u User
	var Ulist []User
	var Tuid []Follow
	db.Table("follows").Where("Uid = ?", Uid).Find(&Tuid)
	for _, v := range Tuid {
		var u User
		db.Table("users").Where("id = ?", v.TUid).First(&u)
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
	db, err := gorm.Open("mysql", "root:123456@(localhost)/douyin?charset=utf8mb4&parseTime=True&Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&Follow{})
	var Ulist []User
	var Uid []Follow
	db.Table("follows").Where("TUid = ?", tuid).Find(&Uid)
	for _, v := range Uid {
		var u User
		db.Table("users").Where("id = ?", v.Uid).First(&u)
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
	db, err := gorm.Open("mysql", "root:123456@(localhost)/douyin?charset=utf8mb4&parseTime=True&Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var Ulist []User
	var Flist []User
	var Tuid []Follow
	db.Table("follows").Where("uid = ?", uid).Find(&Tuid)
	for _, v := range Tuid {
		var u User
		db.Table("users").Where("id = ?", v.TUid).First(&u)
		Ulist = append(Ulist, u)
	}
	var Uid []Follow
	db.Table("follows").Where("TUid = ?", uid).Find(&Uid)
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
