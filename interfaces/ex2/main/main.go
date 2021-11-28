//Пример на соответствие интерфейсу.
//Интерфейс - это контракт.

package main

import "fmt"

type Stream interface {
	read() string
	write(string)
	close()
}

func writeToStream(stream Stream, text string) {
	stream.write(text)
}

func closeStream(stream Stream) {
	stream.close()
}

//структура файл
type File struct {
	text string
}

//структура папка
type Folder struct{}

//реализация методов для типа *File
func (f *File) read() string {
	return f.text
}

//реализация методов для типа *Folder
func (f *File) write(message string) {
	f.text = message
	fmt.Println("Запись в файл строки", message)
}

func (f *File) close() {
	fmt.Println("Файл закрыт")
}

//реализация методов для типа *Folder
// func (f *Folder) read() string {
// 	return "читаем"
// }

// func (f *Folder) write(message string) {
// 	fmt.Println("Запись в папку", message)
// }

func (f *Folder) close() {
	fmt.Println("Закрываем папку ")
}

func main() {
	myFile := &File{} //!!! метод реализован именно для типа *File, то есть указатель на объект File, а не для типа File.
	myFolder := &Folder{}

	writeToStream(myFile, "hello world")
	closeStream(myFile)

	//closeStream(myFolder) //Ошибка: тип *Folder не реализует интерфейс Stream. Т.е не хватает выше закоментированных методов (read и write)
	myFolder.close() //поэтому вызвать close можно только так.
}
