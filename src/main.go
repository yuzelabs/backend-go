package main

import "fmt"

func hello(message string) string {
	if message == "" {
		message = "Hello world"
	}

	return message
}

func add(a, b int) int {
	return a + b
}

func main() {
	sum := add(2, 5)

	fmt.Println(hello("ABC"))
	fmt.Println(sum)
}
