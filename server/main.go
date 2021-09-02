package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var clients Clients

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	ValidMoves map[StdPosn][]StdPosn `json:"validMoves"`
	RecentMove Move                  `json:"lastMove"`
}

func startGame() {
	for idx, client := range clients {
		if idx > 0 {
			err := client.WriteMessage(1, []byte("Starting game."))
			if err != nil {
				log.Println(err)
			}
		}
	}

	game := NewGame()
	for !game.gameIsOver() {
		activePlayer := clients[game.getActivePlayer()]

		// send prelim info over
		msg := Message{game.validMoves(), game.lastMove()}
		jsonMsg, _ := json.MarshalIndent(msg, "", "  ")
		err := activePlayer.WriteMessage(1, jsonMsg)
		if err != nil {
			log.Println(err)
		}

		// read move from activePlayer
		_, p, err := activePlayer.ReadMessage()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(p))

		// game.move(,)
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade connection to a WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")
	clients.push(ws)
	if clients.isFull() {
		startGame()
	}
}

func main() {
	fmt.Println("Server running...")
	http.HandleFunc("/ws", wsEndpoint)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
