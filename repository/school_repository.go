package repository

import (
	"notebook-backend/repository/model"

	"gorm.io/gorm"
)

type SchoolRepository interface {
	FindById(id uint) (*model.School, error)
	FindByUserId(userID uint) ([]model.School, error)
	Create(school model.School) (model.School, error)
}

type schoolRepository struct {
	db *gorm.DB
}

func NewSchoolRepository(db *gorm.DB) SchoolRepository {
	return &schoolRepository{db: db}
}

func (r *schoolRepository) FindById(id uint) (*model.School, error) {
	var school model.School
	err := r.db.Where("id = ?", id).First(&school).Error
	if err != nil {
		return nil, err
	}
	return &school, nil
}

func (r *schoolRepository) FindByUserId(userID uint) ([]model.School, error) {
	var school []model.School
	err := r.db.Where("user_id = ?", userID).Find(&school).Error
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
