package repository

import (
	"github.com/Higins/blocks/domain"
	appModel "github.com/Higins/blocks/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var log = logrus.WithFields(logrus.Fields{"type": "user_repository"})

type UserRepository struct{}

func NewUserRepository() domain.UserRepository {
	return &UserRepository{}
}
func (m *UserRepository) Get(tx *gorm.DB, userId int) (*appModel.User, error) {
	res := appModel.User{}
	err := tx.First(&res, userId).Error
	if err != nil {
		log.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("can't find user")
		return nil, err
	}
	return &res, nil
}
func (m *UserRepository) GetUserByEmail(tx *gorm.DB, email string) (*appModel.User, error) {
	res := appModel.User{}
	err := tx.Where("email = ?", email).First(&res).Error
	if err != nil {
		log.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("can't find user by email")
		return nil, err
	}
	return &res, nil
}
