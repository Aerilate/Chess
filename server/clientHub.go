package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type ClientHub [3]*websocket.Conn

func (c *ClientHub) push(conn *websocket.Conn) {
	if c[1] == nil {
		c[1] = conn
	} else if c[2] == nil {
		c[2] = conn
	}
}

func (c *ClientHub) isFull() bool {
	return c[1] != nil && c[2] != nil
}

func (c *ClientHub) getConn(p Player) *websocket.Conn {
	if p == Player1 {
		return c[1]
	}
	return c[2]
}

func (c *ClientHub) broadCast(message []byte) (err error) {
	if c[1] == nil || c[2] == nil {
		return fmt.Errorf("player not found")
	}
	if err = c[1].WriteMessage(1, message); err != nil {
		return err
	}
	if err = c[2].WriteMessage(1, message); err != nil {
		return err
	}
	return nil
}
