package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var address = flag.String("address", ":8080", "Address to listen on.")

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello from a macOS HTTP service")
	})

	log.Printf("HTTP server listening on %s", *address)
	if err := http.ListenAndServe(*address, nil); err != nil {
		log.Fatal(err)
	}
}
