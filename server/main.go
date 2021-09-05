package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

var clients ClientHub

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type CheckMessage struct {
	Check string `json:"check"`
}

type Message struct {
	ValidMoves map[string][]string `json:"validMoves"`
	Fen        string              `json:"fen"`
}

func startGame() {
	for idx, client := range clients {
		if idx > 0 {
			err := client.WriteMessage(1, []byte(`{"Starting game":[]}`))
			if err != nil {
				log.Println(err)
			}
		}
	}

	game := NewGame()
	for !game.IsOver() {
		activePlayer := clients.getConn(game.ActivePlayer())
		msg := Message{ValidMoves: game.ValidMoves(), Fen: game.Fen()}
		jsonMsg, err := json.MarshalIndent(msg, "", "  ")
		if err != nil {
			log.Println(err)
		}

		err = activePlayer.WriteMessage(1, jsonMsg)
		if err != nil {
			log.Println(err)
		}

		// read move from activePlayer
		_, p, err := activePlayer.ReadMessage()
		if err != nil {
			log.Println(err)
		}
		rank, _ := strconv.ParseInt(string(p[1]), 10, 32)
		src := StdPosn{file: rune(p[0]), rank: int(rank)}
		rank, _ = strconv.ParseInt(string(p[3]), 10, 32)
		dest := StdPosn{file: rune(p[2]), rank: int(rank)}
		move := Move{src, dest}
		fmt.Println(src, dest)
		game.Move(move)

		checkMsg := CheckMessage{string(game.Checked())}
		jsonMsg, err = json.MarshalIndent(checkMsg, "", "  ")
		err = clients.broadcast(jsonMsg)
		if err != nil {
			log.Println(err)
		}
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
