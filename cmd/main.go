package main

import (
	"httpServer/api"
	"net/http"
)

func main() {
	http.Handle("/datetime", http.HandlerFunc(routes.DateTime))
	http.ListenAndServe(":8080", nil)
}
