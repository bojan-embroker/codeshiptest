package dump

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func Dump(responseWriter http.ResponseWriter, request *http.Request) {

	dump, _ := httputil.DumpRequest(request, true)
	log.Print(string(dump))
	responseWriter.Write(dump)
}
