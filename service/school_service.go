package service

import (
	"errors"
	"notebook-backend/handler/dto"
	"notebook-backend/helper"
	"notebook-backend/repository"
	"notebook-backend/repository/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SchoolService interface {
	GetSchoolByUserId(userId string) ([]dto.SchoolResponse, error)
	CreateSchool(schoolInput dto.CreateSchool) (dto.SchoolResponse, error)
}

type schoolService struct {
	schoolRepo repository.SchoolRepository
	userRepo   repository.UserRepository
}

func NewSchoolService(schoolRepo repository.SchoolRepository, userRepo repository.UserRepository) SchoolService {
	return &schoolService{schoolRepo: schoolRepo, userRepo: userRepo}
}

func (s *schoolService) GetSchoolByUserId(userId string) ([]dto.SchoolResponse, error) {
	parsedUUID, err := uuid.Parse(userId)

	user, err := s.userRepo.FindByID(parsedUUID)
	if err != nil {
		return []dto.SchoolResponse{}, errors.New("User not found")
	}

	schools, err := s.schoolRepo.FindByUserId(user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []dto.SchoolResponse{}, nil
		}
		return nil, errors.New("database error")
	}

	schoolMap := []dto.SchoolResponse{}
	for _, school := range schools {
		schoolMap = append(schoolMap, dto.SchoolResponse{
			Name:      school.Name,
			Address:   school.Address,
			Telephone: school.Telephone,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return schoolMap, nil
}

func (s *schoolService) CreateSchool(schoolInput dto.CreateSchool) (dto.SchoolResponse, error) {
	parsedUUID, err := uuid.Parse(schoolInput.UserID)

	user, err := s.userRepo.FindByID(parsedUUID)
	if err != nil {
		return dto.SchoolResponse{}, errors.New("User not found")
	}

	schools, err := s.schoolRepo.FindByUserId(user.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.SchoolResponse{}, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) || !hasDuplicateSchool(schools, schoolInput.Name) {

		newSchool := model.School{
			UserID:    user.ID,
			Name:      schoolInput.Name,
			Address:   schoolInput.Address,
			Telephone: schoolInput.Telephone,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		createdSchool, err := s.schoolRepo.Create(newSchool)
		if err != nil {
			return dto.SchoolResponse{}, helper.ErrInsertRecord
		}

		return dto.SchoolResponse{
			Name:      createdSchool.Name,
			Address:   createdSchool.Address,
			Telephone: createdSchool.Telephone,
			CreatedAt: createdSchool.CreatedAt,
			UpdatedAt: createdSchool.UpdatedAt,
		}, nil
	}
	return dto.SchoolResponse{}, errors.New("duplicate school name")

}

func hasDuplicateSchool(schools []model.School, name string) bool {
	for _, school := range schools {
		if school.Name == name {
			return true
		}
	}
	return false
}
