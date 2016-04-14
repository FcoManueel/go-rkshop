package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func helloFunction(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hola, yo soy una funcion~!")
}

func helloFile(w http.ResponseWriter, req *http.Request) {
	fileContent, err := ioutil.ReadFile("file.txt")
	if err != nil {
		panic(err.Error())
	}
	w.Write(fileContent)
}

func main() {
	log.Println("[go-short] starting...")

	http.HandleFunc("/func", helloFunction)
	http.HandleFunc("/file", helloFile)
	err := http.ListenAndServe(":8080", nil)

	log.Fatal(err)
}
