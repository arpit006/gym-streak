package models

import "time"

type User struct {
	//GymStreakAppModel
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
