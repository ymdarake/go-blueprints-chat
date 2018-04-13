package main

import (
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	socket *websocket.Conn
	send   chan *Message
	room   *room
	// data in cookieUserData
	UserData map[string]interface{}
}

func (c *Client) read() {
	for {
		var msg *Message
		if err := c.socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now()
			msg.Name = c.UserData["name"].(string)
			if avatarURL, ok := c.UserData["avatar_url"]; ok {
				msg.AvatarURL = avatarURL.(string)
			}
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *Client) write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
