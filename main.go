package main

import "fmt"

func main() {
	fmt.Println("hi")

	for nextMessage() {
		handleMessage()
	}
}

func handleMessage(_ any) {}
