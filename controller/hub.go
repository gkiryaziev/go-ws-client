package controller

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"

	"github.com/gkiryaziev/go-ws-client/service"
)

type Hub struct {
	ws        *websocket.Conn
	send      chan []byte
	broadcast chan []byte
	topics    service.TopicPool
	debug     bool
}

// NewHub return new Hub object.
func NewHub(ws *websocket.Conn, t service.TopicPool, debug bool) *Hub {

	return &Hub{
		ws:        ws,
		send:      make(chan []byte, 256),
		broadcast: make(chan []byte),
		topics:    t,
		debug:     debug,
	}
}

// Send data.
func (h *Hub) Send(data []byte) {

	// debug message
	if h.debug {
		log.Println("Sent:", string(data))
	}

	h.send <- data
}

// Reader is data reader.
func (h *Hub) Reader() {
	defer h.ws.Close()
	log.Println("Reader started.")
	for {
		_, message, err := h.ws.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		h.broadcast <- message
	}
}

// Writer is data writer.
func (h *Hub) Writer() {
	defer h.ws.Close()
	log.Println("Writer started.")
	for {
		select {
		case message, ok := <-h.send:
			if ok {
				err := h.ws.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					log.Println(err)
					break
				}
			}
		}
	}
}

// Run is main method.
func (h *Hub) Run() {
	log.Println("Hub started.")
	for {
		select {
		case b := <-h.broadcast:

			// debug message
			if h.debug {
				log.Println("Received:", string(b))
			}

			// unmarshal message
			var msg WSMessage
			err := json.Unmarshal(b, &msg)
			if err != nil {
				log.Println(err)
				break
			}

			// check action
			switch msg.Action {
			case "PUBLISH":
				// get func by topic name
				if fn, ok := h.topics[msg.Topic]; ok {
					if m := fn(msg.Data); m != nil {
						h.Send(m)
						time.Sleep(time.Millisecond * 500)
					}
				}
			}
		}
	}
}
