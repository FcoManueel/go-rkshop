package main

import (
        "log"
)

func swap(a, b int) {
        log.Println("swap!")
        a, b = b, a
}

func main() {
        a := 1
        b := 2

        log.Println("a:", a, " b:", b)

        swap(a, b)
        log.Println("a:", a, " b:", b)
}
