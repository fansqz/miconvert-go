package ws

import (
	"log"
	"net/http"
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
	client := &Client{conn: conn, send: make(chan []byte, 256)}
	CM.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
