package server

import (
	"encoding/json"
	"net/http"
)

type Handler struct{


}


func (h *Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var requestDTO RequestDTO;

	err := json.NewDecoder(req.Body).Decode(&requestDTO)

	if err != nil{
		res.WriteHeader(http.StatusBadRequest);
		res.Write([]byte(err.Error()));
		return;
	}


}


type RequestDTO struct {
	Content string `json:"content"`
}
