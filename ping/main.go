package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// PingHandlerFunc is an api that sends a ping request
func PingHandlerFunc(w http.ResponseWriter, r *http.Request) {
	validToken, err := CreateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	// Attach the token to the header
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:9001", nil)
	if err != nil {
		fmt.Fprintf(w, "unable to make request: %v", err)
	}

	req.Header.Set("Authorization", validToken) // TODO Append Bearer
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "unable to make request: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, "unable to make request: %v", err)
	}

	fmt.Fprint(w, "Ping! \n")
	fmt.Fprintf(w, string(body))

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", PingHandlerFunc)

	http.ListenAndServe(":9000", r)
}
