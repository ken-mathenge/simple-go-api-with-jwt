package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// PongHandlerFunc returns a pong after an authorized ping
func PongHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong!")
}

func main() {
	r := mux.NewRouter()

	r.Use(AuthenticationMiddleware())
	r.HandleFunc("/", PongHandlerFunc)

	http.ListenAndServe(":9001", r)
}
