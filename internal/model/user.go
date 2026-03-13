package model

import "time"

type User struct {
	ID          string
	Coin        int
	LastCheckin time.Time
}
