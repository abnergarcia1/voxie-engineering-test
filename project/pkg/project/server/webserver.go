package server

import (
	"io"
	"net/http"
)

type WebServer struct{}

func (ws *WebServer)WebClient(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/webclient" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}


	io.WriteString(w, "Hello, Mux!\n")
}