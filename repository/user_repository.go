package repository

import (
	"notebook-backend/repository/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	Create(user model.User) (model.User, error)
	FindByID(userId uuid.UUID) (model.User, error)
	FindByUsernameAndPassword(username string, password string) (model.User, error)
	Update(user model.User) (model.User, error)
	Delete(userId uuid.UUID) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Create(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) FindByID(userId uuid.UUID) (model.User, error) {
	var user model.User
	err := r.db.First(&user, "uuid = ?", userId).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) FindByUsernameAndPassword(username string, password string) (model.User, error) {
	var user model.User
	err := r.db.Where(&model.User{Username: username, Password: password}).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) Update(user model.User) (model.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) Delete(userID uuid.UUID) error {
	err := r.db.Where("uuid = ?", userID).Delete(&model.User{}).Error
	if err != nil {
		return err
	}
	return nil
}
