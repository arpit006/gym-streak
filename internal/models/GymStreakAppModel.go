package models

import "time"

type GymStreakAppModel struct {
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	selfSubStruct interface{} // store pointer to the actual sub-struct
}
