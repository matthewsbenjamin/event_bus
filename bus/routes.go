package main

func (s *server) routes() {

	// Just return some documentation
	s.Router.HandleFunc("/", s.log(s.indexHandler()))

}
