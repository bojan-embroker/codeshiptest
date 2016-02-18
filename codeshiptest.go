package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/bojan-embroker/codeshiptest/dump"
)

func main() {

	test := flag.String("test", "", "Blah")
	flag.Parse()

	if *test == "success" {
		log.Println("Success")
		os.Exit(0)
	}

	if *test == "fail" {
		log.Println("Fail")
		os.Exit(-1)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {

		dump.Dump(responseWriter, request)
		log.Println("### test01 ###")
	})

	http.ListenAndServe(":8080", mux)
}
