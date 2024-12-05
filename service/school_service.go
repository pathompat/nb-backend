package service

import (
	"errors"
	"notebook-backend/handler/dto"
	"notebook-backend/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SchoolService interface {
	GetSchoolByUserId(userId string) ([]dto.SchoolResponse, error)
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
