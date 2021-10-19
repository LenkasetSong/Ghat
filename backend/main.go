package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// upgrader: websocket.Upgrader
// Read and Write buffer size required
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// CheckOrigin: checks the origin of the connection
	// allows requests from the frontend (React.js) development server to here
	// temporarily always returns true
	CheckOrigin: func(r *http.Request) bool { return true },
}

// reader: listens for new messages being sent to our websocket endpoint
func reader(conn *websocket.Conn) {
	for {
		// read in a message
		var messageType, p, err = conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out for clarification
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

// serveWs: websocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// upgrade this connection for new messages coming through
	// on our websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming through
	// on our websocket connection
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
	// map `/ws` to `serveWs`
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App v0.0.1")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8081", nil))
}
