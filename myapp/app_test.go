package myapp

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIndexPathHandler(*testing.T) {

	muxtest := http.NewServeMux()
	muxtest.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Wolrd")

	})
	http.ListenAndServe("", muxtest)

	//
}
