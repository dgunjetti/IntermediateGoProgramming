package main

import (
	"fmt"
	"log"
	"net/http"
)

//START OMIT
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//END OMIT
