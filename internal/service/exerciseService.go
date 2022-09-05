package service

import (
	"github.com/arpit006/gym_streak/internal/exceptions"
	"github.com/arpit006/gym_streak/internal/models"
	"github.com/arpit006/gym_streak/internal/stores"
	"github.com/google/uuid"
)

type exerciseServices struct{}

var exerciseService *exerciseServices = nil

var exerciseStore = stores.GetExerciseStoreInstance()

func GetExerciseServiceInstance() *exerciseServices {
	if exerciseService != nil {
		return exerciseService
	}
	exerciseService = &exerciseServices{}
	return exerciseService
}

func (exService exerciseServices) AddAnExercise(exercise *models.Exercise) exceptions.GymStreakAppException {
	// validate Exercise body
	err := validateAddExerciseRequest(exercise)
	if err != nil {
		return err
	}
	// call the store to save to database
	return exerciseStore.AddAnExercise(exercise)
}

func validateAddExerciseRequest(exercise *models.Exercise) exceptions.GymStreakAppException {
	exercise.ExerciseId = uuid.New()
	if exercise.Name == "" {
		return exceptions.NewHttpException(400, exceptions.BAD_REQUEST, "exercise name is blank", "name is blank", nil)
	}
	if exercise.Category == "" {
		return exceptions.NewHttpException(400, exceptions.BAD_REQUEST, "exercise category is blank", "category is blank", nil)
	}
	if exercise.Type == "" {
		return exceptions.NewHttpException(400, exceptions.BAD_REQUEST, "exercise type is blank", "type is blank", nil)
	}
	exerciseTypeError := exercise.Type.IsValidExerciseType()
	if exerciseTypeError != nil {
		return exerciseTypeError
	}
	exerciseCategoryError := exercise.Category.IsValidExerciseCategory()
	if exerciseCategoryError != nil {
		return exerciseCategoryError
	}
	return nil
}
