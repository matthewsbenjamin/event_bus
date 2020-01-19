package main

func (s *server) routes() {

	// Just return some documentation
	// TODO - extract this from a file
	s.Router.HandleFunc("/", s.log(s.handleIndex())).Methods("GET")

	// search the buffer for unconsumed event
	s.Router.HandleFunc("/testing", s.log(s.tempGetAddEvent())).Methods("GET")

	// TODO get all events in the buffer
	s.Router.HandleFunc("/events", s.log(s.handleGetAllEvents())).Methods("GET")

	// TODO get all events in the buffer
	s.Router.HandleFunc("/events/{event}", s.log(s.handleGetSpecificEvent())).Methods("GET")

	// post a new event to the queue
	s.Router.HandleFunc("/", s.log(s.handleNewEvent())).Methods("POST")

	// Post succesfully handled events by some service
	s.Router.HandleFunc("/events/{event}/services/{service}", s.log(s.handleHandledEvent())).Methods("POST")

	// search the buffer for unconsumed event
	s.Router.HandleFunc("/events/{event}/services/{service}", s.log(s.handleGetUnconsumedEvent())).Methods("GET")

}

// Need to be able to:

// Need to handle:
// get requests for some documentation - from a file
// posts of new events - post full event to event buffer/message queue
// post succesful handling of existing events (not as new events)
// get requests searching for unconsumed event(s?) (<n)
//
// Graceful empty returns

// Want:
// get a specific event
// purge all unexpired events (for manual expiry)
