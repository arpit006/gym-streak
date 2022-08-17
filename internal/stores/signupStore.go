package stores

import (
	"fmt"
	"github.com/arpit006/gym_streak/internal/logger"
	"github.com/arpit006/gym_streak/internal/models"
)

type SignupStore struct{}

var signupStore *SignupStore = nil

// GetSignupStoreInstance returns singleton instance of SignupStore
func GetSignupStoreInstance() *SignupStore {
	if signupStore != nil {
		return signupStore
	}
	signupStore = &SignupStore{}
	return signupStore
}

func (ss *SignupStore) Signup(user models.User) error {
	query := "INSERT INTO USERS(id, first_name, last_name) VALUES(?,?,?)"
	_, err := DB.db.Query(query, user.ID, user.FirstName, user.LastName)
	if err != nil {
		logger.ErrorF(fmt.Sprintf("Error while saving signup info. Error is %s", err))
		return err
	}
	logger.InfoF("User signed-up successfully")
	return nil
}
