package main

import (
	"net/http"

	"github.com/bojan-embroker/codeshiptest/dump"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {

		dump.Dump(responseWriter, request)
	})

	http.ListenAndServe(":8080", mux)
}
