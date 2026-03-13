package handler

import (
	"encoding/json"
	"net/http"

	"overlix-backend/internal/service"
)

type CheckinRequest struct {
	UserID string `json:"userId"`
}

func DailyCheckin(w http.ResponseWriter, r *http.Request) {

	var req CheckinRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	user, err := service.DailyCheckin(req.UserID)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "check-in success",
		"coin":    user.Coin,
	})
}
