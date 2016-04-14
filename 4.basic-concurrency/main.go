package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	boring("boring!")
	//    go boring("async boring!")
	fmt.Println("end")
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(2 * time.Second)
	}
}
