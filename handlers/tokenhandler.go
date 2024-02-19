package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sanjaysingh/aad-jwt-gen/services"
)

type TokenRequest struct {
	ClientId     string `json:"clientId"`
	TenantId     string `json:"tenantId"`
	ClientSecret string `json:"clientSecret"`
	Scope        string `json:"scope"`
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req TokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := services.GenerateToken(req.ClientId, req.TenantId, req.ClientSecret, req.Scope)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
