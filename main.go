package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func printInitial(w http.ResponseWriter, r *http.Request) {
	fmt.Println("initial commit")
}

// PORT is the default port
const PORT = "8080"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", printInitial)
	fmt.Println("Now listening on port: ", PORT)
	log.Fatal(http.ListenAndServe(":"+ PORT, r))
}
