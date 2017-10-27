package main

import (
	"fmt"
	"github.com/gucosme/an-api/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.DefaultHandler)
	fmt.Println("Listening at 3005")
	http.ListenAndServe(":3005", nil)
}
