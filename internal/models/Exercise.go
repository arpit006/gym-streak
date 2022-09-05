package models

import (
	"encoding/json"
	"github.com/arpit006/gym_streak/internal/constants"
	"github.com/arpit006/gym_streak/internal/logger"
	"github.com/google/uuid"
)

type Exercise struct {
	ExerciseId uuid.UUID                  `json:"id"`
	Name       string                     `json:"name"`
	MetaInfo   string                     `json:"meta"`
	Category   constants.ExerciseCategory `json:"category"`
	Type       constants.ExerciseType     `json:"type"`
	//BodyPart
	//	List of imageIds
}

func (e *Exercise) ToString() string {
	s, err := json.Marshal(e)
	if err != nil {
		logger.Errorf("Error converting Exercise obj to string. Error is %s", err.Error())
		return ""
	}
	return string(s)
}
