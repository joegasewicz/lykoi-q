package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var FLAGS = []string{"-p", "-t", "--debug"}

func GetIntFromArg(args []string, flag string, index int) (int, error) {
	var err error
	if len(args) == index {
		return 0, errors.New(fmt.Sprintf("missing %s argument value", flag))
	}
	argInt, err := strconv.Atoi(args[index+1])
	if err != nil {
		return 0, errors.New(fmt.Sprintf("didn't pass a correct argument to %s", flag))
	}
	return argInt, err
}

func main() {
	config := Config{}
	server := QueueServer{}
	args := os.Args

	config.Process(args)
	logging := NewLogging(config.Level, 1)
	logging.Log(fmt.Sprintf("Starting server on %d", config.Port), "test")

	http.HandleFunc("/enqueue", server.HandleProvider)
	http.HandleFunc("/dequeue", server.HandleConsumer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
