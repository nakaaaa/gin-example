package gateway

import (
	"errors"

	"github.com/Nakaaaa/gin-example/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (rp *UserRepository) GetByUserID(db *gorm.DB, userID int) (*model.User, error) {
	var model *model.User

	err := db.Where("`user_id` = ?", userID).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
			return nil, err
		}

		return nil, err
	}

	return model, nil
}

func (rp *UserRepository) AddOrUpdate(db *gorm.DB, user *model.User) (*model.User, error) {
	err := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(user).Error
	if err != nil {
		return nil, err
	}

	return rp.GetByUserID(db, user.UserID)
}

func (rp *UserRepository) DeleteUser(db *gorm.DB, userID int) error {
	err := db.Where("`user_id` = ?", userID).Delete(&model.User{}).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	return nil
}
