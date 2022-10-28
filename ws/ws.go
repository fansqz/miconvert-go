package ws

import (
	"github.com/gorilla/websocket"
	"log"
	"miconvert-go/utils"
	"net/http"
	"strings"
	"time"
)

//
// ServeWs
//  @Description: http升级到ws，并添加client
//  @param hub
//  @param w
//  @param r
//
func ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//获取token
	array := strings.Split(r.URL.String(), ",")
	token := array[len(array)-1]
	user, _ := utils.ParseToken(token)
	WSManager.wsMap[user.Id] = conn

	go listenClose(user.Id, conn)
}

// 监听连接是否断开
func listenClose(userid int, ws *websocket.Conn) {
	defer func() {
		WSManager.DeleteConn(userid)
	}()
	ws.SetReadLimit(maxMessageSize)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
	}
}
