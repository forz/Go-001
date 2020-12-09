package main

import (
	"Go-000/Week03/cycle"
	"Go-000/Week03/signal"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hi")) // nolint
}

func NewServer1() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", Hello)
	return &http.Server{
		Addr:    "0.0.0.0:8081",
		Handler: mux,
	}
}
func NewServer2() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", Hello)
	return &http.Server{
		Addr:    "0.0.0.0:8082",
		Handler: mux,
	}
}

func NewServer3() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", Hello)
	return &http.Server{
		Addr:    "0.0.0.0:8083",
		Handler: mux,
	}
}
func main() {
	var cycle cycle.Cycle
	server1 := NewServer1()
	server2 := NewServer2()
	server3 := NewServer3()
	cycle.AddServer(server1)
	cycle.AddServer(server2)
	cycle.AddServer(server3)
	signal.HookSignals(&cycle)
	err := cycle.Run()
	if err != nil {
		log.Println("main.error:", err)
	}
}
