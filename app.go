package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func serveApp(addr string) func(http.ResponseWriter, *http.Request) {
	target, err := url.Parse(addr)
	if err != nil {
		log.Fatalf("Unable to parse app server addr: %v (%v)", addr, err)
	}
	proxy := httputil.NewSingleHostReverseProxy(target)
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}
