package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func WSHandler(ws *websocket.Conn) {
	for {
		var message string

		// Receive message from the client
		err := websocket.Message.Receive(ws, &message)
		if err != nil {
			fmt.Println("error receiving message:", err)
			return
		}
		fmt.Println("Message from client:", message)

		// Send a response back to the client
		response := "Response from server: " + message
		err = websocket.Message.Send(ws, response)
		if err != nil {
			fmt.Println("error sending message:", err)
			return
		}
	}
}

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", SayHello)

	handler.Handle("/ws", websocket.Handler(WSHandler))

	http.ListenAndServe(":3000", handler)

}





