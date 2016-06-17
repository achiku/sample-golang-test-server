package client

import (
	"fmt"
	"log"
	"net/http"
)

func hello(logger Logfer) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Logf("yeaaaaaaaahhh!!")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "hello!")
	}
}

func goodbye(logger Logfer) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Logf("goodby!!!!!")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "goodbye!")
	}
}

type handler func(w http.ResponseWriter, r *http.Request)

// URLHandlerMap url and handler
type URLHandlerMap []map[string]handler

// NewMockServerMux creates new mock server mux
func NewMockServerMux(hm URLHandlerMap, logger Logfer) *http.ServeMux {
	mux := http.NewServeMux()
	if hm == nil {
		mux.HandleFunc("/hello", hello(logger))
		mux.HandleFunc("/goodbye", goodbye(logger))
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
	logger := mockServerLogger{}
	mux := NewMockServerMux(nil, logger)
	server := &http.Server{
		Handler: mux,
		Addr:    "localhost:" + port,
	}
	return server
}

// Logfer inerface Logf
type Logfer interface {
	Logf(format string, args ...interface{})
}

type mockServerLogger struct{}

func (l mockServerLogger) Logf(format string, args ...interface{}) {
	log.Println(format, args)
}
