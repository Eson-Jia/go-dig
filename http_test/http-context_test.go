package main

import (
	"log"
	"net/http"
	"testing"
	"time"
)

func TestHttpContext(t *testing.T) {
	http.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		select {
		case <-req.Context().Done():
			log.Println("cancel")
		case <-time.After(time.Second * 5):
			w.Write([]byte("pong\n"))
		}
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}