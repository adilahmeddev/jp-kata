package adapter

import (
	"encoding/json"
	"fmt"
	jpservice "github.com/adilahmeddev/jp-kata/backend/pkg/domain/jp-service"
	"net/http"
	"strings"
)

type Handler struct {
	Processor jpservice.LanguageProcessor
}

func NewHandler(processor jpservice.LanguageProcessor) *Handler {
	return &Handler{Processor: processor}
}

func (h *Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var requestDTO RequestDTO

	err := json.NewDecoder(req.Body).Decode(&requestDTO)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}
	fmt.Println("decoded body, got :" + requestDTO.Content)
	uniqueCharacters := h.Processor.GetUnique(requestDTO.Content)

	successDTO := SuccessDTO{
		Success: true,
		Data:    uniqueCharacters,
	}
	fmt.Println("encodidng response body, got :" + strings.Join(uniqueCharacters, ""))

	err = json.NewEncoder(res).Encode(successDTO)

	if err != nil {
		failureDTO := FailureDTO{
			Success: false,
			Data:    nil,
		}
		res.WriteHeader(501)
		json.NewEncoder(res).Encode(failureDTO)
	}
}

type RequestDTO struct {
	Content string `json:"content"`
}

type SuccessDTO struct {
	Success bool     `json:"success"`
	Data    []string `json:"data"`
}

type FailureDTO struct {
	Success bool     `json:"success"`
	Data    []string `json:"data"`
}
