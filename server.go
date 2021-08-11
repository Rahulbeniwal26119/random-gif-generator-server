package main

import (
	"fmt"
	"log"
	"net/http"
	"random_gif/gif_generator"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/gif", generate_gif)
	http.HandleFunc("/", root_handler)
	http.HandleFunc("/count", count_request)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func generate_gif(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count += 1
	gif_generator.GifGenerator(w)
	mu.Unlock()
}
func root_handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "<b>Visit to <i>/gif</i> to generate gif</b><br>")
	fmt.Fprintf(w, "<b>Visit to <i>/count</i> to  count the requests</b>")
	count += 1
	mu.Unlock()
}
func count_request(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
