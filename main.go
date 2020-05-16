package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// PORT is the default port
const PORT = "8080"

var addr = flag.String("addr", "127.0.0.1:8080", "http service address")

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	startServer()
}

func getBuildPath() string {
	buildPath, exists := os.LookupEnv("BUILD_PATH")

	if exists {
		return buildPath
	}
	log.Panic("No static/build path found")
	return ""
}

// NewRouter is a new gorilla/mux router
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(getBuildPath())))
	return r
}

func startServer() {
	r := NewRouter()
	fmt.Println("Now listening on port: ", PORT)
	if err := http.ListenAndServe(":"+PORT, r); err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
