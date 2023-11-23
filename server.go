package main

import (
	"fmt"
	"net/http"
)

type Options struct {
}

type QueueServer struct {
	Options *Options
}

// HandleProvider
func (s *QueueServer) HandleProvider(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Povider...")
}

// HandleConsumer
func (s *QueueServer) HandleConsumer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "consumer...")
}
