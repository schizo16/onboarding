package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/invite", generateInvite)
	http.HandleFunc("/invite/", resolveInvite)

	http.ListenAndServe(":8080", nil)
}

type Invite struct {
	Code       string
	WorldID    string
	HostUserID string
	ExpireAt   time.Time
	UsageLimit int
	UsedCount  int
}

var inviteStore = map[string]*Invite{}

func generateCode() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func generateInvite(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		WorldID    string `json:"worldId"`
		HostUserID string `json:"hostUserId"`
		Mode       string `json:"mode"` // single | unlimited
	}

	var req Request
	json.NewDecoder(r.Body).Decode(&req)

	code := generateCode()

	limit := 1
	if req.Mode == "unlimited" {
		limit = -1
	}

	invite := &Invite{
		Code:       code,
		WorldID:    req.WorldID,
		HostUserID: req.HostUserID,
		ExpireAt:   time.Now().Add(24 * time.Hour),
		UsageLimit: limit,
	}

	inviteStore[code] = invite

	resp := map[string]string{
		"link": "overlix.net/invite/" + code,
	}

	json.NewEncoder(w).Encode(resp)
}

func resolveInvite(w http.ResponseWriter, r *http.Request) {

	code := strings.TrimPrefix(r.URL.Path, "/invite/")

	invite, ok := inviteStore[code]
	if !ok {
		http.Error(w, "invalid code", 404)
		return
	}

	if time.Now().After(invite.ExpireAt) {
		http.Error(w, "expired", 400)
		return
	}

	if invite.UsageLimit != -1 &&
		invite.UsedCount >= invite.UsageLimit {
		http.Error(w, "used", 400)
		return
	}

	invite.UsedCount++

	worldInfo := map[string]string{
		"worldId":    invite.WorldID,
		"hostUserId": invite.HostUserID,
	}

	json.NewEncoder(w).Encode(worldInfo)
}
