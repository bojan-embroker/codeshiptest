package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/bojan-embroker/codeshiptest/dump"
)

func main() {

	finish := flag.Bool("finish", false, "")
	flag.Parse()

	if *finish {
		log.Println("\"finish\" flag supplied, finishing execution")
	} else {
		log.Println("running server...")
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {

			dump.Dump(responseWriter, request)
		})

		http.ListenAndServe(":8080", mux)
	}
}
