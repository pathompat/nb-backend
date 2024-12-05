package repository

import (
	"notebook-backend/repository/model"

	"gorm.io/gorm"
)

type SchoolRepository interface {
	FindByUserId(userId uint) ([]model.School, error)
	Create(school model.School) (model.School, error)
}

type schoolRepository struct {
	db *gorm.DB
}

func NewSchoolRepository(db *gorm.DB) SchoolRepository {
	return &schoolRepository{db: db}
}

func (r *schoolRepository) FindByUserId(userId uint) ([]model.School, error) {
	var school []model.School
	err := r.db.Unscoped().Where("user_id = ?", userId).Find(&school).Error
	if err != nil {
		return []model.School{}, err
	}
	return school, nil
}

func (r *schoolRepository) Create(school model.School) (model.School, error) {
	err := r.db.Unscoped().Create(&school).Error
	if err != nil {
		return model.School{}, err
	}
	return school, nil
}
