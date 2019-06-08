package main

import (
	"fmt"
	"log"
	"net/http"
)

//START OMIT
type Server struct{}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
func main() {
	s := &Server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//END OMIT
