package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Event is the only way that a service can communicate with the bus
// One (or more) of these will be returned in JSON representation
// by the get requests
type Event struct {
	Notification string
	PostedBy     string
	PostedOn     time.Time
	Expiry       int64
	Payload      map[string]string
	ConsumedBy   map[string]int
}

type server struct {
	EventBuffer *map[string]Event
	Router      *mux.Router
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.Router.ServeHTTP(w, req)
}

func newServer() *server {
	s := &server{}
	s.Router = mux.NewRouter()
	s.routes()

	e := make(map[string]Event)
	s.EventBuffer = &e

	return s
}

func main() {

	// Event buffer

	S := newServer()

	// SERVER
	http.ListenAndServe(":8080", S)

}
