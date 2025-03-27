// Copyright (c) Jeff Mendoza <jlm@jlm.name>
// SPDX-License-Identifier: MIT

package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello/", helloHandler).Methods("GET")
	router.Use(func(h http.Handler) http.Handler {
		return handlers.LoggingHandler(os.Stdout, h)
	})
	http.ListenAndServe(":7000", router)
}
