package constants

import (
	"fmt"
	"github.com/arpit006/gym_streak/internal/exceptions"
)

// ExerciseType [cardio, strength]
type ExerciseType string

var (
	CARDIO_EXERCISE_TYPE            = newExerciseType("cardio")
	STRENGTH_TRAINING_EXERCISE_TYPE = newExerciseType("strength_training")
)

func newExerciseType(exType string) ExerciseType {
	return ExerciseType(exType)
}

func GetStringValueOfExerciseType(exType ExerciseType) string {
	return string(exType)
}

func (exType ExerciseType) IsValidExerciseType() exceptions.GymStreakAppException {
	switch exType {
	case CARDIO_EXERCISE_TYPE, STRENGTH_TRAINING_EXERCISE_TYPE:
		return nil
	}
	return exceptions.NewHttpException(400, exceptions.BAD_REQUEST, fmt.Sprintf("invalid exercise type [%s]", exType), "invalid type", nil)
}
