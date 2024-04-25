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
		msg := scanner.Bytes()
        method, content, err := rpc.DecodeMessage(msg)
        if err != nil {
            logger.Printf("Got an error: %s", err)
        }
		scanner.Split(rpc.Split)
		handleMessage(logger,msg)
	}
}

func handleMessage(logger *log.Logger,msg  any) {
    logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey, you didnt give me a good file")
	}
	return log.New(logfile, "[educationalsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
