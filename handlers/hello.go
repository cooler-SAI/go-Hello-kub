package handlers

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintln(w, "Hello World")
	if err != nil {
		return
	}
}
