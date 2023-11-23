package main

import (
	"fmt"
	"net/http"
	"os"
)

var FLAGS = []string{"-p", "-t", "--debug"}

func main() {
	config := Config{}
	server := QueueServer{}
	args := os.Args
	queue := QueueInit(false)

	config.Process(args)

	logging := NewLogging(config.Level, 1)
	logging.Log(fmt.Sprintf("Starting server on %d", config.Port), "test")

	http.HandleFunc("/enqueue", server.HandleProvider(queue, logging))
	http.HandleFunc("/dequeue", server.HandleConsumer)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
