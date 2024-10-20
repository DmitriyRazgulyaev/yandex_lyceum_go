package main

import (
	"fmt"
)

func main() {
	a := []byte("str")
	fmt.Println(utf8.(a))
}

var names = make(map[string]string)

func PrintHello(name string) string {
	names[name] = name
	return fmt.Sprintf("Hello, %s!", name)
}
