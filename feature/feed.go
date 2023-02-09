package feature

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Feed(C *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@(localhost)/douyin?charset=utf8mb4&parseTime=True&Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&Video{}, &User{})
	var VL []Video

	db.Table("videos").Find(&VL)
	for i, V := range VL {
		key := V.UID
		var u User
		db.Table("users").Where("v_key = ?", key).Last(&u)
		VL[i].Author = u
		// fmt.Println(u)
		// fmt.Println(VL[i])
	}
	C.JSON(http.StatusOK, Video_Feed{
		StatusCode: 0,
		StatusMsg:  "string",
		NextTime:   time.Now().Unix(),
		Vlist:      VL,
	})
}
