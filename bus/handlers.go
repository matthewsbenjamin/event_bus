package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gomarkdown/markdown"
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

		fmt.Fprintln(w, http.StatusNoContent)

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

// func (s *server) handleGetSpecificEvent() http.HandlerFunc {

// 	return func(w http.ResponseWriter, req *http.Request) {

// 		fmt.Fprintln(w, "TODO One event")

// 	}

// }

// func (s *server) handleHandledEvent() http.HandlerFunc {

// 	return func(w http.ResponseWriter, req *http.Request) {

// 		fmt.Fprintln(w, "TODO Handle handled event")

// 	}

// }

// // searching the buffer for unconsumed events
// // can only return ONE event at a time - allows for service competition
// func (s *server) handleGetUnconsumedEvent() http.HandlerFunc {

// 	return func(w http.ResponseWriter, req *http.Request) {

// 		fmt.Fprintln(w, "TODO Unconsumed Event")

// 	}

// }
