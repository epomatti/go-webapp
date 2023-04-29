package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/healthcheck", healthcheck)

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", "dev")
	fmt.Fprintf(w, "version: %s\n", "1.0.0")
}
