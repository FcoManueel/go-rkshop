package main

import (
	"io"
	"log"
	"net/http"
)

func helloFunction(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hola, yo soy una funcion~!")
}

func main() {
	log.Println("[go-short] starting...")

	http.HandleFunc("/func", helloFunction)
	err := http.ListenAndServe(":8080", nil)

	log.Fatal(err)
}

//func helloFile(w http.ResponseWriter, req *http.Request) {
//	fileContent, err := ioutil.ReadFile("file.txt")
//	if err != nil {
//		panic(err.Error())
//	}
//	w.Write(fileContent)
//}

//http.HandleFunc("/file", helloFile)
