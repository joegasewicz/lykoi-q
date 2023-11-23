package main

type ProviderReq struct {
	Message string `json:"message"`
}

type ProviderRes struct {
	ID        uint32
	CreatedOn string
}
