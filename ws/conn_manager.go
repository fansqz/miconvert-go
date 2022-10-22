package ws

import (
	"github.com/gorilla/websocket"
	"sync"
)

type wsManager struct {
	wsMap map[int]*websocket.Conn
	mutex sync.Mutex
}

func (wm *wsManager) sendMessage(userId int, bytes []byte) {
	wm.mutex.Lock()
	conn := wm.wsMap[userId]
	w, err := conn.NextWriter(websocket.TextMessage)
	if err != nil {
		wm.deleteConn(userId)
	}
	w.Write(bytes)
}

func (wm *wsManager) deleteConn(userId int) {
	ws := wm.wsMap[userId]
	ws.Close()
	WSManager.mutex.Lock()
	delete(WSManager.wsMap, userId)
	WSManager.mutex.Unlock()
}

func (wm *wsManager) addConn(userId int, ws *websocket.Conn) {
	wm.mutex.Lock()
	wm.wsMap[userId] = ws
	wm.mutex.Unlock()
}

var WSManager = &wsManager{
	wsMap: make(map[int]*websocket.Conn),
}
