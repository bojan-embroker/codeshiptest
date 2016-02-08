package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {

		dump, _ := httputil.DumpRequest(request, true)
		log.Print(string(dump))
		responseWriter.Write(dump)
	})

	http.ListenAndServe(":8080", mux)
}
