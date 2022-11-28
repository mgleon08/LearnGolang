package main

import (
    "fmt"
    "learnpackage/custompackage"
)

func main() {
    text := custompackage.SayHi("Foo")
    fmt.Println(text)
}
