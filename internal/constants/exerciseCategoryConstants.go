package constants

import (
	"fmt"
	"github.com/arpit006/gym_streak/internal/exceptions"
)

// ExerciseCategory [isometric, compound]
type ExerciseCategory string

// supported categories for exercises
var (
	ISOMETRIC_EXERCISE_CATEGORY = newExerciseCategory("isometric")
	COMPOUND_EXERCISE_CATEGORY  = newExerciseCategory("compound")
)

func newExerciseCategory(category string) ExerciseCategory {
	return ExerciseCategory(category)
}

func (exCategory ExerciseCategory) GetStringValueOfExerciseCategory() string {
	return string(exCategory)
}

func (exCategory ExerciseCategory) IsValidExerciseCategory() exceptions.GymStreakAppException {
	switch exCategory {
	case ISOMETRIC_EXERCISE_CATEGORY, COMPOUND_EXERCISE_CATEGORY:
		return nil
	}
	return exceptions.NewHttpException(400, exceptions.BAD_REQUEST, fmt.Sprintf("invalid exercise category [%s]", exCategory), "invalid category", nil)
}
