package main

import (
	"log"
)

type Age int

func main() {
	var number int = 5
	var age Age = 6

	log.Println("number:", number, "age:", age)

	number = age // ...?
}
