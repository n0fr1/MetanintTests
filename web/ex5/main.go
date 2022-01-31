package main

import (
	"fmt"
	"net/http"
)

type helloHandler struct {
	subject string
}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", h.subject)
}

func main() {

	worldHandler := &helloHandler{"World"}
	roomHandler := &helloHandler{"Mark"}
	undefined := &helloHandler{"undefined path"}

	http.Handle("/world", worldHandler)
	http.Handle("/room", roomHandler)
	http.Handle("/", undefined)

	http.ListenAndServe(":8000", nil)
}
