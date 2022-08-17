package models

import (
	"encoding/json"
	"github.com/arpit006/gym_streak/internal/logger"
	"time"
)

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (u User) ToString() string {
	s, err := json.Marshal(u)
	if err != nil {
		logger.ErrorF("Could not Marshall User object to JSON")
		return ""
	}
	return string(s)
}
