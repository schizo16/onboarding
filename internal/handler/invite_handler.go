package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"overlix-backend/internal/service"
)

type CreateInviteRequest struct {
	WorldID    string `json:"worldId"`
	HostUserID string `json:"hostUserId"`
	Mode       string `json:"mode"`
}

func GenerateInvite(w http.ResponseWriter, r *http.Request) {

	var req CreateInviteRequest
	json.NewDecoder(r.Body).Decode(&req)

	invite := service.CreateInvite(
		req.WorldID,
		req.HostUserID,
		req.Mode,
	)

	json.NewEncoder(w).Encode(map[string]string{
		"link": "https://overlix.net/invite/" + invite.Code,
	})
}

func ResolveInvite(w http.ResponseWriter, r *http.Request) {

	code := strings.TrimPrefix(r.URL.Path, "/invite/")

	invite, err := service.ResolveInvite(code)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"worldId":    invite.WorldID,
		"hostUserId": invite.HostUserID,
	})
}
