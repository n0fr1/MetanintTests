//1. curl localhost:8000 -X POST -d '{"name":"John", "age":30, "salary":3000.50}' - провалимся в default: unknown content type
//2. curl localhost:8000 -X POST -d '{"name":"John", "age":30, "salary":3000.50}' -H "Content-Type: application/json" - провалимся в application/json
//3. curl localhost:8000 -X POST -d '{"name":"John", "age":30, "salary":3000.50}' -H "Content-Type: application/xml" - провалимся в application/xml

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Employee struct {
	Name   string  `json:"name" xml:"name"`
	Age    int     `json:"age" xml:"age"`
	Salary float32 `json:"salary" xml:"salary"`
}

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

	var employee Employee

	contentType := r.Header.Get("Content-Type")

	switch contentType {

	case "application/json":

		err := json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
			return
		}

	case "application/xml":

		err := json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			http.Error(w, "Unable to unmarshal XML", http.StatusBadRequest)
			return
		}

	default:
		http.Error(w, "Unknown content type", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Got a new employee!\nName: %s\nAge: %dy.o.\nSalary %0.2f\n",
		employee.Name,
		employee.Age,
		employee.Salary,
	)
}
