package websocket

import (
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

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	var ws, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ws, nil
}

/*
// Reader: listens for new messages being sent to our websocket endpoint
func Reader(conn *websocket.Conn) {
	for {
		fmt.Println("Reader: Reading")
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out for clarification
		// fmt.Println(string(p))
		// echo
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

// Writer: echo
func Writer(conn *websocket.Conn) {
	for {
		fmt.Println("Writer: Sending")
		messageType, r, err := conn.NextReader()
		if err != nil {
			fmt.Println(err)
			return
		}
		w, err := conn.NextWriter(messageType)
		if err != nil {
			fmt.Println(err)
			return
		}
		if _, err := io.Copy(w, r); err != nil {
			fmt.Println(err)
			return
		}
		if err := w.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}
}
*/
