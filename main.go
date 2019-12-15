package main

import (
	"app/handle"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle.Handle(handle.Home))
	http.HandleFunc("/create", handle.Handle(handle.Create))
	http.ListenAndServe(":8080", nil)
}
