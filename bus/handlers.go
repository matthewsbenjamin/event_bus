package main

import (
	"fmt"
	"net/http"
)

func (s *server) indexHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		fmt.Fprintln(w, "Documentation will eventually go here on how to communicate with this bus")

	}

}
