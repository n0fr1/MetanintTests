//Проверка GET-запроса: http://localhost:8000/?name=John%20%D0%BC
//Проверка POST-запроса: curl localhost:8000 -X POST -d 'Hello,World!'

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Handler struct {
}

func main() {

	handler := &Handler{}
	http.Handle("/", handler)

	srv := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	srv.ListenAndServe()

}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet: //перехватываем GET-запрос.
		name := r.FormValue("name")
		fmt.Fprintf(w, "Parsed query-param with key \"name\": %s", name)
	case http.MethodPost: //перехватываем POST-запрос.
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		fmt.Fprintf(w, "Parsed request body: %s\n", string(body))
	}
}
