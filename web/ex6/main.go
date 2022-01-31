package main

import (
	"fmt"
	"net/http"
	"time"
)

type helloHandler struct {
	subject string
}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", h.subject)
}

func main() {

	roomHandler := &helloHandler{"Mark"}
	http.Handle("/room", roomHandler)

	srv := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	srv.ListenAndServe()
}
