package main

import (
	"fmt"
	"ghat/backend/pkg/websocket"
	"log"
	"net/http"
)

// serveWs: websocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.Host)
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		// write to http.ResponseWriter
		fmt.Fprintf(w, "%+V\n", err)
	}
	// 2 threads
	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	// map `/ws` to function `serveWs`
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App v0.0.2")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8081", nil))
}
