package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {


	for _, e := range os.Environ() {
        pair := strings.Split(e, "=")
	fmt.Fprintf(w, "\n", r.URL.Path[1:])
	fmt.Fprintf(w, pair[0], r.URL.Path[1:])}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server Running...")
	http.ListenAndServe(":8080", nil)
}

