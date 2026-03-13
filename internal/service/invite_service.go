package service

import (
	"errors"
	"math/rand"
	"time"

	"overlix-backend/internal/model"
)

var inviteStore = map[string]*model.Invite{}

func generateCode() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func CreateInvite(worldID, hostUserID, mode string) *model.Invite {

	limit := 1
	if mode == "unlimited" {
		limit = -1
	}

	invite := &model.Invite{
		Code:       generateCode(),
		WorldID:    worldID,
		HostUserID: hostUserID,
		ExpireAt:   time.Now().Add(24 * time.Hour),
		UsageLimit: limit,
	}

	inviteStore[invite.Code] = invite
	return invite
}

func ResolveInvite(code string) (*model.Invite, error) {

	invite, ok := inviteStore[code]
	if !ok {
		return nil, errors.New("invalid code")
	}

	if time.Now().After(invite.ExpireAt) {
		return nil, errors.New("expired")
	}

	if invite.UsageLimit != -1 &&
		invite.UsedCount >= invite.UsageLimit {
		return nil, errors.New("used")
	}

	invite.UsedCount++
	return invite, nil
}
