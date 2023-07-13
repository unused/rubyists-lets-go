package main

import (
	"fmt"
	"net/http"
)

// Server provides a counter that is increased with every visit.
type Server struct {
	Count int
}

// ServeHTTP is the expected interface method to make server an HTTP handler.
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.Count = s.Count + 1
	fmt.Fprintf(w, "%d\n", s.Count)
}

func main() {
	handler := Server{0}
	http.ListenAndServe(":6301", &handler)
}
