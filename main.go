package main

import (
	"fmt"

	go_say_hello "github.com/umarbawazirfaiz/go-say-hello/v2"
)

func main() {
	sayHello := go_say_hello.SayHello("Umar")

	fmt.Println(sayHello)
}
