package service

import (
	"errors"
	"time"

	"overlix-backend/internal/model"
)

var userStore = map[string]*model.User{}

func DailyCheckin(userID string) (*model.User, error) {

	user, ok := userStore[userID]
	if !ok {
		user = &model.User{
			ID:   userID,
			Coin: 0,
		}
		userStore[userID] = user
	}

	now := time.Now()

	// đã checkin hôm nay?
	if sameDay(user.LastCheckin, now) {
		return nil, errors.New("already checked in today")
	}

	reward := 100
	user.Coin += reward
	user.LastCheckin = now

	return user, nil
}

func sameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
