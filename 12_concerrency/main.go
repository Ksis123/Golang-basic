package main

import "fmt"

func hello(s string) {
	for {
		fmt.Println(s)
	}
}

func main() {
	go hello("Hello from another thread 1")

	fmt.Println("Hello from main thread")
}