package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get")
	})

	mux.HandleFunc("POST /v1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("post")
		w.Write([]byte("GET /v1/user/{username}"))
		// w.WriteHeader(200)
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		panic(err)
	}

}
