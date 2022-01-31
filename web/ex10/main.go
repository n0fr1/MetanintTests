//загрузка файлов на сервер
//curl localhost:8000 -F 'file=@testfile.txt' http://localhost/home/evgeniy/upload
//curl -i http://127.0.0.1:8000 - тестируем сервер

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type UploadHandler struct {
	UploadDir string
}

func main() {

	uploadHandler := &UploadHandler{
		UploadDir: "upload",
	}

	http.Handle("/upload", uploadHandler)

	srv := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	srv.ListenAndServe()

}

func (h *UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	file, header, err := r.FormFile("file") //file - io.Reader - для потокового считывания файла. header - различная доп.информация о файле. err - если r.FormFile - не смог считать файл из переданной формы
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file) //считываем содержимое файла в переменную data []byte для последующей записи на диск
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
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
