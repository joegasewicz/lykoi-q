package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type QueueServer struct {
}

// HandleProvider

func (s *QueueServer) HandleProvider(queue *Queue, logging *Logging) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var data ProviderReq
		var message interface{}
		var createdOn time.Time
		var jsonRes ProviderRes
		var jsonData []byte
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error: povider JSON missing 'message' property\n")
			return
		}
		err = json.Unmarshal(body, &data)
		message = data.Message
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error: povider JSON is malformed or missing\n")
			return
		}
		createdOn = time.Now()
		err = queue.Enqueue(&message)
		if err != nil {
			goto Errors
		}
		w.Header().Set("Content-Type", "application/json")
		jsonRes = ProviderRes{
			ID:        queue.LinkedList.ListTail().ID,
			CreatedOn: createdOn,
		}
		jsonData, err = json.Marshal(jsonRes)
		if err != nil {
			goto Errors
		}
		_, err = fmt.Fprintf(w, string(jsonData))
		if err != nil {
			goto Errors
		}
		logging.Log("Messaged ID:", strconv.Itoa(jsonRes.ID))
		return
	Errors:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Server error\n%e", err)
		return

	}

}

// HandleConsumer
func (s *QueueServer) HandleConsumer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "consumer...")
}
