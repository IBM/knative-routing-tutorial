package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"math/rand"
	"time"
)

var words = []string{"fantastic", "wonderful", "super"}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world received a request.")
    rand.Seed(time.Now().UnixNano())
	fmt.Fprintf(w, "Hello and have a %s day!\n", words[rand.Intn(len(words))])
}

func main() {
	log.Print("Hello world sample started.")

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
