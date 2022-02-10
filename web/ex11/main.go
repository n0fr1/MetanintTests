//посмотреть файлы на сервере:
//localhost:8080/testfile.txt

package main

import (
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

	//file server
	dirToServe := http.Dir(uploadHandler.UploadDir)

	fs := &http.Server{
		Addr:         ":8080",
		Handler:      http.FileServer(dirToServe),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fs.ListenAndServe()
}
