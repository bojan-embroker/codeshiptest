package main

import (
	"flag"
	"fmt"
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
		responseWriter.Write([]byte(fmt.Sprintf("Process ID: %v", os.Getpid())))
	})

	mux.HandleFunc("/crash", func(responseWriter http.ResponseWriter, request *http.Request) {
		os.Exit(1)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, mux)
}
