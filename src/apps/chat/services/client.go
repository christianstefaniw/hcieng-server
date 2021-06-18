package services

import (
	"context"
	"encoding/json"
	accounts "hciengserver/src/apps/account/services"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type client struct {
	room   *Room
	conn   *websocket.Conn
	msg    chan *message
	ctx    context.Context
	cancel context.CancelFunc
	*accounts.Account
}

const (
	readWait       = 5 * time.Second
	writeWait      = 5 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

func ServeWs(r *Room, user *accounts.Account, conn *websocket.Conn) {
	ctx, cancel := context.WithCancel(context.Background())
	c := &client{room: r, conn: conn, msg: make(chan *message, 256), ctx: ctx, cancel: cancel, Account: user}
	c.room.register <- c
	go c.doWork(r.AdminTextOnly)
}

func (c *client) unregister() {
	close(c.msg)
	c.room.unregister <- c
	c.conn.Close()
}

func (c *client) Cancel() {
	c.cancel()
}

func (c *client) read(errChan chan error) {
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(_ string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				errChan <- err
			}
			c.Cancel()
		} else {
			c.room.broadcast <- msgStringToMessage(msg, c.Account)
		}
	}
}

func (c *client) write(errChan chan error) {
	ticker := time.NewTicker(pingPeriod)

	for {
		select {
		case msg := <-c.msg:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				errChan <- err
			}

			json.NewEncoder(w).Encode(msg)

			if qued := len(c.msg); qued > 0 {
				for i := 0; i < qued; i++ {
					json.NewEncoder(w).Encode(<-c.msg)
				}
			}

			if err := w.Close(); err != nil {
				c.cancel()
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				errChan <- err
			}
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *client) doWork(adminTextOnly bool) {
	errChan := make(chan error)
	go c.read(errChan)

	if !adminTextOnly || c.Admin {
		go c.write(errChan)
	}

	for {
		select {
		case <-c.ctx.Done():
			c.unregister()
			log.Printf("%s client closed", c.EmailAddr)
			return
		case <-c.room.ctx.Done():
			c.unregister()
			log.Printf("%s room closed", c.room.Id)
			return
		case err := <-errChan:
			log.Printf("client error: %s", err)
		}
	}
}
