package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/echo", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)

		_, _ = writer.Write([]byte(`{text:"Hello world"}`))

		fmt.Println(writer.Header())
	})
	log.Fatalln(http.ListenAndServe("localhost:8080", nil))
}
