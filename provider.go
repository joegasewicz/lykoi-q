package main

import "time"

type ProviderReq struct {
	Message string `json:"message"`
}

type ProviderRes struct {
	ID        int       `json:"id"`
	CreatedOn time.Time `json:"created_on"`
}
