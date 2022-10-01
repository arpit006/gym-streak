package stores

import (
	"fmt"
	"github.com/arpit006/gym_streak/internal/exceptions"
	"github.com/arpit006/gym_streak/internal/models"
	"github.com/arpit006/gym_streak/internal/query"
)

type exerciseStores struct{}

var exerciseStore *exerciseStores = nil

var exerciseQuery = query.GetExerciseQueryInstance()

func GetExerciseStoreInstance() *exerciseStores {
	if exerciseStore != nil {
		return exerciseStore
	}
	exerciseStore = &exerciseStores{}
	return exerciseStore
}

func (eStore exerciseStores) AddAnExercise(e *models.Exercise) exceptions.GymStreakAppException {
	insertQuery := exerciseQuery.GetCreateExerciseQuery()
	_, err := STORE.db.Exec(insertQuery, e.ExerciseId, e.Name, e.MetaInfo, e.Category, e.Type)
	if err != nil {
		return exceptions.NewDatabaseException(500, exceptions.DATABASE_SAVE_ERROR, fmt.Sprintf("Error while saving data in database. Error is %s", err), "error in saving data to database", nil)
	}
	return nil
}

func (eStore exerciseStores) UpdateAnExercise(e models.Exercise) {

}

func (eStore exerciseStores) DeleteAnExercise(e models.Exercise) {

}

func (eStore exerciseStores) GetAllExercise() {

}
