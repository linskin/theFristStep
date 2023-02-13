package feature

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var MessageId = int64(1)

var M_count sync.Map
var Num_Message sync.Map

func MessageAction(c *gin.Context) {

	DB.AutoMigrate(&Message{})
	a_t := c.Query("action_type")
	text := c.Query("content")
	Token := c.Query("token")
	var u User
	DB.Table("users").Where("token = ?", Token).First(&u)
	S_Tuid := c.Query("to_user_id")
	Tuid, _ := strconv.Atoi(S_Tuid)
	K := GetKey(u.ID, Tuid)
	DB.Model(&Message{}).Count(&MessageId)
	//MessageId++
	if a_t == "1" {
		atomic.AddInt64(&MessageId, 1)
		m := Message{
			ID:           MessageId,
			Content:      text,
			From_user_id: int64(u.ID),
			To_user_id:   int64(Tuid),
			CreateTime:   time.Now().Local().Unix(), //time.Now().Format("2006-01-02 15:04:05"),
			MKey:         K,
		}
		DB.Model(&Message{}).Create(&m)
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "success",
	})
}

func MessageChat(c *gin.Context) {
	DB.AutoMigrate(&Message{})
	Token := c.Query("token")
	S_Tuid := c.Query("to_user_id")
	var u User
	DB.Table("users").Where("token = ?", Token).First(&u)
	Tuid, _ := strconv.Atoi(S_Tuid)
	K := GetKey(u.ID, Tuid)
	K1 := fmt.Sprintf("%d_%d", Tuid, u.ID)
	var MList []Message
	var None_List []Message
	num, _ := M_count.Load(K1)
	var count int
	if num.(int) == 0 {
		DB.Model(&Message{}).Where("M_key = ?", K).Find(&MList)
		DB.Model(&Message{}).Where("M_key = ? ", K).Count(&count)
		Num_Message.Store(K1, count)
		fmt.Println(MList)
		c.JSON(http.StatusOK, MessageResponse{
			StatusCode:  0,
			MessageList: MList,
		})
		M_count.Delete(K1)
		M_count.Store(K1, 1)
	} else {
		DB.Model(&Message{}).Where("M_key = ? ", K).Count(&count)
		DB.Model(&Message{}).Where("M_key = ? ", K).Last(&MList)
		v, _ := Num_Message.Load(K1)
		if v.(int) != count && MList[0].To_user_id == int64(u.ID) {
			Num_Message.Delete(K1)
			Num_Message.Store(K1, count)
			fmt.Println(MList)
			c.JSON(http.StatusOK, MessageResponse{
				StatusCode:  0,
				MessageList: MList,
			})
		} else {
			c.JSON(http.StatusOK, MessageResponse{
				StatusCode:  0,
				MessageList: None_List,
			})
		}
	}
}

func GetKey(id1 int, id2 int) string {
	var K string
	if id2 > id1 {
		K = strconv.Itoa(id2) + "_" + strconv.Itoa(id1)
	} else {
		K = strconv.Itoa(id1) + "_" + strconv.Itoa(id2)
	}
	return K
}
