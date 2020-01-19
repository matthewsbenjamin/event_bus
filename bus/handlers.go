package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (s *server) handleIndex() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		fmt.Fprintln(w, "Index - Documentation will eventually go here on how to communicate with this bus")

	}

}

func (s *server) handleGetAllEvents() http.HandlerFunc {

	var q []Event

	return func(w http.ResponseWriter, req *http.Request) {

		json.NewDecoder(req.Body).Decode(&q)

		fmt.Fprintln(w, q)

	}

}

func (s *server) handleGetSpecificEvent() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		fmt.Fprintln(w, "TODO One event")

	}

}

func (s *server) handleNewEvent() http.HandlerFunc {

	file, _ := ioutil.ReadFile("temp.json")

	// Marshal that file into the event buffer

	_ = json.Unmarshal(file, &s.EventBuffer)

	return func(w http.ResponseWriter, req *http.Request) {

		fmt.Fprintln(w, "TODO - handle new event")

	}

}

func (s *server) tempGetAddEvent() http.HandlerFunc {

	file, _ := ioutil.ReadFile("temp.json")

	fmt.Println(file)

	// Marshal that file into the event buffer

	_ = json.Unmarshal(file, &s.EventBuffer)

	return func(w http.ResponseWriter, req *http.Request) {

		fmt.Fprintln(w, "TODO - handle new event")

	}

}

func (s *server) handleHandledEvent() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		fmt.Fprintln(w, "TODO Handle handled event")

	}

}

// searching the buffer for unconsumed events
// can only return ONE event at a time - allows for service competition
func (s *server) handleGetUnconsumedEvent() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		fmt.Fprintln(w, "TODO Unconsumed Event")

	}

}
