package main

import (
	"github.com/adilahmeddev/jp-kata/backend/pkg/adapter"
	jpservice "github.com/adilahmeddev/jp-kata/backend/pkg/domain/jp-service"
	"log"
	"net/http"
)

func main() {
	http.Handle("/unique-tokens", &adapter.Handler{
		Processor: &jpservice.JpProcessor{},
	})
	err := http.ListenAndServe(":7290", nil)
	if err != nil {
		log.Fatal(err)

	}
}
