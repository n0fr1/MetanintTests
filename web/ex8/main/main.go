//Проверка GET-запроса: http://localhost:8000/?name=John%20%D0%BC
//Проверка POST-запроса: curl localhost:8000 -X POST -d 'Hello,World!'
//Проверка POST-запроса: curl localhost:8000 -X POST -d '{"name":"John", "age":30, "salary":3000.50}'
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Employee struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Salary float32 `json:"salary"`
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

	switch r.Method {

	case http.MethodGet: //перехватываем GET-запрос.

		name := r.FormValue("name")
		fmt.Fprintf(w, "Parsed query-param with key \"name\": %s", name)

	case http.MethodPost: //перехватываем POST-запрос.

		// body, err := ioutil.ReadAll(r.Body) //первый вариант
		// if err != nil {
		// 	http.Error(w, "Unable to parse request body", http.StatusBadRequest)
		// 	return
		// }
		// defer r.Body.Close()

		// var employee Employee
		// err = json.Unmarshal(body, &employee)
		// if err != nil {
		// 	http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
		// 	return
		// }

		var employee Employee
		err := json.NewDecoder(r.Body).Decode(&employee) //второй вариант. Передаем напрямую без использование буфера - body.

		if err != nil {
			http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Got a new employee!\nName: %s\nAge: %dy.o.\nSalary %0.2f\n",
			employee.Name,
			employee.Age,
			employee.Salary,
		)

	}
}
