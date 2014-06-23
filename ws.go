package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const (
	pingPeriod = 60 * time.Second // XXX this should be a cmd line flag
	writeWait  = 10 * time.Second // XXX this too
)

var (
	pool     *redis.Pool
	upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
)

type connection struct {
	ws           *websocket.Conn
	subscription string
	send         chan []byte
}

func (c *connection) readFromRedis() {
	conn := pool.Get()
	defer conn.Close()

	psc := redis.PubSubConn{conn}
	if err := psc.Subscribe(c.subscription); err != nil {
		log.Fatalf("Failed to subscribe to %v: %v", c.subscription, err)
		return
	}
	log.Printf("Connected to redis channel %v", c.subscription)
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			log.Printf("Got a redis message: %v", v)
			c.send <- v.Data
		case redis.Subscription:
			log.Print("Got a redis subscription")
			// XXX nop?
		case error:
			log.Fatalf("Error reading messages: %v", v)
		default:
			log.Fatalf("Got an unknown redis message type: %v", v)
		}
	}
}

func (c *connection) writeMessages() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()

	for {
		log.Print("Awaiting things to write")
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				log.Printf("Error receiving message to write")
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				log.Printf("Fatal error sending message: %v", err)
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				log.Printf("Fatal error sending ping: %v", err)
				return
			}
		}

	}
}

func (c *connection) write(messageType int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(messageType, payload)
}

func serveWs(server string, prefix string) func(http.ResponseWriter, *http.Request) {
	pool = newPool(server)
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("ws connection request")
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", 405)
			return
		}
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			if _, ok := err.(websocket.HandshakeError); !ok {
				log.Println(err)
			}
			return
		}
		vars := mux.Vars(r)
		channel := vars["channel"]
		sub := subChannel(prefix, channel)
		c := &connection{
			ws:           ws,
			subscription: sub,
			send:         make(chan []byte, 256),
		}
		go c.writeMessages()
		c.readFromRedis()
	}
}

//
// Based on pool example in redigo documentation
//
func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func subChannel(prefix, path string) string {
	return prefix + path
}
