package main

import (
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {

	server := &PlayerServer{&InMemoryPlayerStore{}}

	// handler := http.HandlerFunc(server.ServeHTTP)

	http.ListenAndServe(":8080", server)
}
