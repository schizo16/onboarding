package model

import "time"

type Invite struct {
	Code       string
	WorldID    string
	HostUserID string
	ExpireAt   time.Time
	UsageLimit int
	UsedCount  int
}
