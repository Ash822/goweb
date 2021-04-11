package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	const port = ":8000"

	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		msg := "Request served the server"
		payload := []byte(`{"message": "` + msg + `"}`)

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write(payload)
	})

	log.Printf("Server is listening at port%s", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
