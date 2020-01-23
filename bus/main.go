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
	EventType  string            `json:"EventType"`
	PostedBy   string            `json:"Posted_by"`
	PostedOn   time.Time         `json:"Posted_on"`
	Expiry     int               `json:"Expiry"`
	Payload    map[string]string `json:"Payload"`
	ConsumedBy map[string]int    `json:"Consumed_by"`
}

type server struct {
	EventBuffer []*Event
	Router      *mux.Router
}

func newServer() *server {

	s := &server{}

	s.Router = mux.NewRouter()
	s.routes()

	e := []*Event{}
	s.EventBuffer = e

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.Router.ServeHTTP(w, req)
}

func main() {

	// Event buffer

	S := newServer()

	// SERVER
	http.ListenAndServe(":8080", S)

}
