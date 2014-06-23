package main

// This implementation owes a lot to the Gorilla websocket chat example
// Copyright 2013 The Gorilla WebSocket Authors.

import (
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	addr         = flag.String("addr", ":8080", "http service address")
	staticDir    = flag.String("static", ".", "static file root")
	staticUrl    = flag.String("static-url", "/static", "static file base url")
	redisServer  = flag.String("redis-server", ":6379", "redis server address")
	channelPfx   = flag.String("channel-prefix", "gawrsh-", "redis channel prefix")
	appProxyAddr = flag.String("proxy-addr", "http://localhost:8000/", "app server proxy address")
)

func main() {
	flag.Parse()

	r := mux.NewRouter()
	r.Handle("/static/{path:.+}", http.FileServer(http.Dir(*staticDir)))
	r.HandleFunc("/ws/{channel}", serveWs(*redisServer, *channelPfx))
	r.HandleFunc("/{path:.*}", serveApp(*appProxyAddr))
	http.Handle("/", r)

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("Unable to start server: ", err)
	}

}
