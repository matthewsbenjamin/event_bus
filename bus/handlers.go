package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gomarkdown/markdown"
	"github.com/gorilla/mux"
)

func (s *server) handleIndex() http.HandlerFunc {

	// Access the documentation
	file, err := ioutil.ReadFile("documentation.md")
	if err != nil {
		return func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintln(w, http.StatusInternalServerError)
		}
	}

	// Convert the markdown file to HTML
	html := markdown.ToHTML(file, nil, nil)

	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, string(html))
	}
}

// hangs off the /events path and returns the entire event queue
func (s *server) handleGetAllEvents() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		json.NewEncoder(w).Encode(&s.EventBuffer)

	}

}

// Temporarily reads the test file and adds to the event queue
func (s *server) tempGetAddEvent() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		file, _ := ioutil.ReadFile("temp.json")

		e := Event{}
		// Marshal that file into the event buffer
		err := json.Unmarshal(file, &e)
		if err != nil {
			fmt.Fprintln(w, http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		// Something here to validate the requests?
		// or is that too much coupled logic...?

		// E is the full queue - this is then appended to
		// TODO LOCK

		s.EventBuffer = append(s.EventBuffer, &e)

		fmt.Fprintln(w, http.StatusCreated)

	}

}

func (s *server) handleNewEvent() http.HandlerFunc {

	file, _ := ioutil.ReadFile("temp.json")

	// Marshal that file into the event buffer
	_ = json.Unmarshal(file, &s.EventBuffer)

	return func(w http.ResponseWriter, req *http.Request) {

		fmt.Fprint(w, "TODO - handle new event")

	}

}

func (s *server) handleGetSpecificEvent() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		id := vars["uid"]

		for _, e := range s.EventBuffer {
			if e.UID == id {

				json.NewEncoder(w).Encode(e)
				return

			}
		}

		fmt.Fprint(w, http.StatusNoContent)
	}
}

func (s *server) handleGetEventType() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		t := vars["type"]
		fmt.Println(t)

		for _, e := range s.EventBuffer {
			if e.EventType == t {
				fmt.Println(e)
				json.NewEncoder(w).Encode(e)
				return

			}
		}

		fmt.Fprint(w, http.StatusNoContent)
	}

}

func (s *server) handleGetEventTypeWithConsumption() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		t := vars["type"]
		serv := vars["service"]
		n, err := strconv.Atoi(vars["n"])
		if err != nil {
			n = 1
		}

		for _, e := range s.EventBuffer {
			if e.EventType == t {

				// found the event, now search the consumed services

				count, ok := e.ConsumedBy[serv]
				fmt.Println(count)
				fmt.Println(ok)
				if ok && count < n {
					json.NewEncoder(w).Encode(e)
					return
				}
				continue
			}
		}

		fmt.Fprint(w, http.StatusNoContent)
	}

}

func (s *server) handleHandledEvent() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		fmt.Fprint(w, "TODO Handle handled event")

	}

}

// // searching the buffer for unconsumed events
// // can only return ONE event at a time - allows for service competition
// func (s *server) handleGetUnconsumedEvent() http.HandlerFunc {

// 	return func(w http.ResponseWriter, req *http.Request) {

// 		fmt.Fprintln(w, "TODO Unconsumed Event")

// 	}

// }
