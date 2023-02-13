package feature

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Feed(C *gin.Context) {
	DB.AutoMigrate(&Video{}, &User{})
	var VL []Video
	DB.Table("videos").Find(&VL)
	for i, V := range VL {
		key := V.UID
		var u User
		DB.Table("users").Where("v_key = ?", key).Last(&u)
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
