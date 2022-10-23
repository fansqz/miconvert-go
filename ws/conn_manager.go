package ws

import (
	"github.com/gorilla/websocket"
	"sync"
)

// todo:直接加锁效率低
type wsManager struct {
	wsMap map[int]*websocket.Conn
	mutex sync.Mutex
}

func (wm *wsManager) SendMessage(userId int, bytes []byte) {
	wm.mutex.Lock()
	conn := wm.wsMap[userId]
	wm.mutex.Unlock()
	if conn == nil {
		return
	}
	w, err := conn.NextWriter(websocket.TextMessage)
	if err != nil {
		wm.DeleteConn(userId)
		return
	}
	w.Write(bytes)
}

func (wm *wsManager) DeleteConn(userId int) {
	WSManager.mutex.Lock()
	ws := wm.wsMap[userId]
	ws.Close()
	delete(WSManager.wsMap, userId)
	WSManager.mutex.Unlock()
}

func (wm *wsManager) AddConn(userId int, ws *websocket.Conn) {
	wm.mutex.Lock()
	wm.wsMap[userId] = ws
	wm.mutex.Unlock()
}

var WSManager = &wsManager{
	wsMap: make(map[int]*websocket.Conn),
}
