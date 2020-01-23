package main

import (
	"net/http"
)

func (s *server) log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// set up some logging functionality in here - make it capture metrics n stuff

		// fmt.Println("Before")
		// defer fmt.Println("After")

		// Then the thing
		next.ServeHTTP(w, r) // call original

	}
}
