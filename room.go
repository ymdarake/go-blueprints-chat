package main

import (
	"log"
	"net/http"

	"github.com/stretchr/objx"

	"github.com/gorilla/websocket"
	"github.com/ymdarake/go-blueprints-tracer/tracer"
)

type room struct {
	// for message
	forward chan *Message
	// for joining Clients
	join chan *Client
	// for leaving Clients
	leave chan *Client
	// Clients in room
	clients map[*Client]bool
	// impl to get avatar url
	getAvatar GetAvatar
	// tracer receives manipulation logs on the room
	Tracer tracer.Tracer
}

func (r *room) Run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			r.Tracer.Trace("a new client joinned")
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			r.Tracer.Trace("a client left")
		case msg := <-r.forward:
			r.Tracer.Trace("received a new message: ", msg.Message)
			for client := range r.clients {
				select {
				case client.send <- msg:
					r.Tracer.Trace(" -- sent to a client")
				default:
					delete(r.clients, client)
					close(client.send)
					r.Tracer.Trace(" -- failed to send to a client")
				}
			}
		}
	}
}

const (
	socketBuffersize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBuffersize, WriteBufferSize: socketBuffersize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	authCookie, err := req.Cookie("auth")
	if err != nil {
		log.Fatal("failed to get cookie: ", err)
		return
	}

	client := &Client{
		socket:   socket,
		send:     make(chan *Message, messageBufferSize),
		room:     r,
		UserData: objx.MustFromBase64(authCookie.Value),
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}

func NewRoom(getAvatar GetAvatar) *room {
	return &room{
		forward:   make(chan *Message),
		join:      make(chan *Client),
		leave:     make(chan *Client),
		clients:   make(map[*Client]bool),
		getAvatar: getAvatar,
		Tracer:    tracer.Off(),
	}
}
