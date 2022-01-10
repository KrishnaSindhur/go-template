package handler

import (
	"net/http"

	"github.com/KrishnaSindhur/go-template/pkg/lib/encode"
)

type healthResponse struct {
	Status string `json:"status"`
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := healthResponse{Status: "ok"}
	encode.WriteJSON(w, res)
}
