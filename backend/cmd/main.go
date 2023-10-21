package main

import (
	"log"

	httpserver "github.com/adilahmeddev/jp-kata/backend/pkg/adapter"
)

func main() {
	server := httpserver.NewServer("7000")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)

	}
}
