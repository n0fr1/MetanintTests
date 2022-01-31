package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) { //назначаем анонимную функцию-обработчик для запросов к серверу, URL которых удовлетворяет маршруту "/hello"
		fmt.Fprintf(w, "Hello World!")
	})

	http.ListenAndServe("localhost:8000", nil) //запускает http-сервер, который будет принимать и обрабатывать запросы на 8000 порту
}
