package serve

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"sync"

	"example.com/m/v2/feature"
)

var K_CMap sync.Map //用来判断用户是否离线，并且存储每个对话窗口对应的conn

func MassageServe() {
	listen, err := net.Listen("tcp", "192.168.1.102:9090")
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go Process(conn)
	}
}

func Process(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf[:]) //读取到buf中
		if err == io.EOF {
			fmt.Println("远程客户端已经退出了")
			return
		}
		var event feature.SendMessageEvent
		err = json.Unmarshal(buf[:n], &event)
		if err != nil {
			panic(err)
		}
		fmt.Println("收到消息:", event)

		FromK := fmt.Sprintf("%d_%d", event.UID, event.TUid)
		if len(event.Content) == 0 {
			K_CMap.Store(FromK, conn)
			continue
		}

		ToK := fmt.Sprintf("%d_%d", event.TUid, event.UID)
		writeConn, exist := K_CMap.Load(ToK) //这力做出来的回传消息conn是any类型的，所以后面要进行断言转换
		if !exist {
			fmt.Printf("用户%d已经离线", event.TUid)
			continue
		}

		pushEvent := feature.PushMessageEvent{
			FromUserId: int64(event.UID),
			MsgContent: event.Content,
		}
		data, err := json.Marshal(pushEvent)
		if err != nil {
			panic(err)
		}
		_, err = writeConn.(net.Conn).Write(data)
		if err != nil {
			panic(err)
		}
	}

}
