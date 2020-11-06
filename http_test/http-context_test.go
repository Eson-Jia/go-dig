package main

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestHttpContext(t *testing.T) {
	http.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		select {
		case <-req.Context().Done():
			log.Println("cancel")
		case <-time.After(time.Second * 5):
			w.Write([]byte("pong"))
			io.Copy(w,req.Body)
		}
	})
	svr:= http.Server{
		Addr: ":8000",
		Handler: nil,
	}
	go func() {
		if err:= svr.ListenAndServe();err!=nil{
			if err == http.ErrServerClosed{
				t.Log("shut down")
			}else{
				t.Fatal(err)
			}
		}
	}()
	resp,err:= http.Post("http://localhost:8000/ping","json",bytes.NewBufferString(`{"msg":"hello"}`))
	if err!=nil{
		t.Fatal(err)
	}
	io.Copy(os.Stdout,resp.Body)
	svr.Shutdown(context.TODO())
}
