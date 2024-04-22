package main

import (
	"io"
	"log"
	"net/http"
)

func bodyclose() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	// good: resp.Body.Close()
	_ = body
}
