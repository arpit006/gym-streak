package stores

import (
	"database/sql"
	"fmt"
	"github.com/arpit006/gym_streak/internal/logger"
	"github.com/arpit006/gym_streak/internal/models"
)

type SignupStore struct {
	db *sql.DB
}

func NewSignupStore(db *sql.DB) *SignupStore {
	return &SignupStore{
		db: db,
	}
}

func (sr SignupStore) Save(user models.User) {
	query := "INSERT INTO USERS(id, first_name, last_name) VALUES(?, ?, ?)"
	logger.PrintAnything(user)
	_, err := sr.db.Query(query, user.ID, user.FirstName, user.LastName)
	if err != nil {
		logger.Error(fmt.Sprintf("Error occurred in signup call.  Error is %s", err))
	}
}
