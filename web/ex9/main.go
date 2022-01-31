//1. curl localhost:8000 -X POST -d '{"name":"John", "age":30, "salary":3000.50}' - провалимся в default: unknown content type
//2. curl localhost:8000 -X POST -d '{"name":"John", "age":30, "salary":3000.50}' -H "Content-Type: application/json" - провалимся в application/json
//3. curl localhost:8000 -X POST -d '{"name":"John", "age":30, "salary":3000.50}' -H "Content-Type: application/xml" - провалимся в application/xml

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

//add
type UploadHandler struct {
	UploadDir string
}

//add

func main() {

	handler := &Handler{}
	http.Handle("/", handler)

	//add
	uploadHandler := &UploadHandler{
		UploadDir: "upload",
	}

	http.Handle("/upload", uploadHandler)

	//add

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

//add
func (h *UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	file, header, err := r.FormFile("file") //file - io.Reader - для потокового считывания файла. header - различная доп.информация о файле. err - если r.FormFile - не смог считать файл из переданной формы
	if err != nil {
		http.Error(w, "Unable to read file - 1", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file) //считываем содержимое файла в переменную data []byte для последующей записи на диск
	if err != nil {
		http.Error(w, "Unable to read file - 2", http.StatusBadRequest)
		return
	}

	filePath := h.UploadDir + "/" + header.Filename //задаем путь, который будет иметь сохраненный на диск файл.

	err = ioutil.WriteFile(filePath, data, 0777) //пишем данные на диск, 1 - путь до файла, 2 - слайс байт, 3 - доступ к заданному файлу
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "File %s has been successfully uploaded", header.Filename) //выводим сообщение об успешно загруженном файле
}

//add
