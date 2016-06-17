package client

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("accessed from client!")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello!")
}

func goodbye(w http.ResponseWriter, r *http.Request) {
	log.Println("accessed from client!")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "goodbye!")
}

// URLHandlerMap url and handler
type URLHandlerMap []map[string]func(w http.ResponseWriter, r *http.Request)

// NewMockServerMux creates new mock server mux
func NewMockServerMux(hm URLHandlerMap) *http.ServeMux {
	mux := http.NewServeMux()
	if hm == nil {
		mux.HandleFunc("/hello", hello)
		mux.HandleFunc("/goodbye", goodbye)
		return mux
	}

	for _, m := range hm {
		for url, handler := range m {
			mux.HandleFunc(url, handler)
		}
	}
	return mux
}

// NewMockServer creates new mock server
func NewMockServer(port string) *http.Server {
	mux := NewMockServerMux(nil)
	server := &http.Server{
		Handler: mux,
		Addr:    "localhost:" + port,
	}
	return server
}
