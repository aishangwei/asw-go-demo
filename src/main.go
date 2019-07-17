package main

import (
	"fmt"
	"net/http"
	"os"
)

func Hello(response http.ResponseWriter, request *http.Request) {
	host,err := os.Hostname()
	if err != nil {
		fmt.Fprintf(response, "Error: %s", err)
	}else {
		fmt.Fprintf(response, "Hostname: %s", host)
	}
}


func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}