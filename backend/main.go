package main

import (
	"fmt"
	"ghat/backend/pkg/websocket"
	"log"
	"net/http"
)

// serveWs: websocket endpoint
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	var conn, err = websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	var client = &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	// map `/ws` to function `serveWs`
	var pool = websocket.NewPool()
	go pool.Start()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Ghat Backend v1.0.0")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8081", nil))
}
