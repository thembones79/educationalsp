package main

import (
	"bufio"
    "educationalsp/rpc"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hi")

    scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
        msg := scanner.Text()
        scanner.Split(rpc.Split)
		handleMessage(msg)
	}
}

func handleMessage(_ any) {}
