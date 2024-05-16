package main

import "fmt"

func Hello(a string) string {
	return "Hello, " + a
}

func main() {
	fmt.Println(Hello("Denis"))
}
