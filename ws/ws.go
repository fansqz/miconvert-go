package ws

import (
	"github.com/gorilla/websocket"
	"log"
	"miconvert-go/utils"
	"net/http"
	"strings"
	"time"
)

var wsMap = make(map[int]*websocket.Conn)

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
	wsMap[user.ID] = conn

	go listenClose(user.ID, conn)
}

func listenClose(userid int, ws *websocket.Conn) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		ws.Close()
		delete(wsMap, userid)
	}()
	for {
		select {
		case <-ticker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
