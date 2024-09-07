package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	const HelloOTUS = "Hello, OTUS!"
	fmt.Println(reverse.String(HelloOTUS))
}
