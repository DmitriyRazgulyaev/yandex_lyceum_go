package main

import (
	"fmt"
)

func main() {

}

var names = make(map[string]string)

func PrintHello(name string) string {
	names[name] = name
	return fmt.Sprintf("Hello, %s!", name)
}
