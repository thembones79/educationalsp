package main

import (
	"bufio"
	"educationalsp/rpc"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("hi")

	logger := getLogger("/home/michu/Projects/go/educationalsp/log.txt")
	logger.Print("Hey, I started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		scanner.Split(rpc.Split)
		handleMessage(msg)
	}
}

func handleMessage(_ any) {}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey, you didnt give me a good file")
	}
	return log.New(logfile, "[educationalsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
