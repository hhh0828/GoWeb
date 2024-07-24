package main

import (
	"net/http"

	"GOWEB/myapp"
)

func main() {
	http.ListenAndServe("", myapp.NewHttpHandler())
}

//
