package main

import (
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryPlayerStore) RecordWin(name string) {

}

func main() {

	server := &PlayerServer{&InMemoryPlayerStore{}}

	// handler := http.HandlerFunc(server.ServeHTTP)

	http.ListenAndServe(":8080", server)
}
