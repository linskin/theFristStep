package feature

import (
	"example.com/m/v2/conf"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Feed(C *gin.Context) {
	conf.DB.AutoMigrate(&Video{}, &User{})
	var VL []Video
	conf.DB.Table("videos").Find(&VL)
	for i, V := range VL {
		key := V.UID
		var u User
		conf.DB.Table("users").Where("v_key = ?", key).Last(&u)
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
