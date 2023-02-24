package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, you've rqeuested: %s\n", request.URL.Path)
	})

	log.Println("Web server is started | localhost:8082")
	http.ListenAndServe(":8082", nil)
}
