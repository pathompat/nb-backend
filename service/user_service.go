package service

import (
	"errors"
	"notebook-backend/handler/dto"
	"notebook-backend/helper"
	"notebook-backend/repository"
	"notebook-backend/repository/model"
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	GetAllUsers() ([]dto.UserResponse, error)
	GetInfoUser(userID string) (dto.UserResponse, error)
	CreateUser(input dto.CreateUser) (dto.UserResponse, error)
	UpdateUser(userID string, input dto.UpdateUser) (dto.UserResponse, error)
	DeleteUser(userID string) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers() ([]dto.UserResponse, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, errors.New("database error")
	}

	userMap := []dto.UserResponse{}
	for _, user := range users {
		userMap = append(userMap, dto.UserResponse{
			UserId:    user.UserId,
			Username:  user.Username,
			StoreName: user.StoreName,
			TierID:    user.TierID,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return userMap, nil
}

func (s *userService) GetInfoUser(userID string) (dto.UserResponse, error) {
	parsedUUID, err := uuid.Parse(userID)
	if err != nil {
		return dto.UserResponse{}, errors.New("invalid UUID format")
	}

	user, err := s.repo.FindByID(parsedUUID)
	if err != nil {
		return dto.UserResponse{}, errors.New("User not found")
	}

	return dto.UserResponse{
		UserId:    user.UserId,
		TierID:    user.TierID,
		Username:  user.Username,
		StoreName: user.StoreName,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *userService) CreateUser(input dto.CreateUser) (dto.UserResponse, error) {
	hashPassword, err := helper.HashPassword(input.Password)
	if err != nil {
		return dto.UserResponse{}, helper.ErrHashPassword
	}

	newUser := model.User{
		Username:  input.Username,
		TierID:    input.TierID,
		StoreName: input.StoreName,
		Password:  hashPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdUser, err := s.repo.Create(newUser)
	if err != nil {
		return dto.UserResponse{}, helper.ErrInsertRecord
	}

	return dto.UserResponse{
		UserId:    createdUser.UserId,
		TierID:    createdUser.TierID,
		Username:  createdUser.Username,
		StoreName: createdUser.StoreName,
		Role:      createdUser.Role,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
	}, nil
}

func (s *userService) UpdateUser(userID string, input dto.UpdateUser) (dto.UserResponse, error) {
	parsedUUID, err := uuid.Parse(userID)
	if err != nil {
		return dto.UserResponse{}, errors.New("invalid UUID format")
	}

	user, err := s.repo.FindByID(parsedUUID)
	if err != nil {
		return dto.UserResponse{}, errors.New("User not found")
	}

	hashPassword, err := helper.HashPassword(input.Password)
	if err != nil {
		return dto.UserResponse{}, helper.ErrHashPassword
	}

	user.Username = input.Username
	user.Password = hashPassword
	user.TierID = input.TierID
	user.StoreName = input.StoreName
	user.UpdatedAt = time.Now()

	updateUser, err := s.repo.Update(user)
	if err != nil {
		return dto.UserResponse{}, errors.New("Failed to update user")
	}

	return dto.UserResponse{
		UserId:    updateUser.UserId,
		TierID:    updateUser.TierID,
		Username:  updateUser.Username,
		StoreName: updateUser.StoreName,
		Role:      updateUser.Role,
		CreatedAt: updateUser.CreatedAt,
		UpdatedAt: updateUser.UpdatedAt,
	}, nil
}

func (s *userService) DeleteUser(userID string) error {
	parsedUUID, err := uuid.Parse(userID)
	if err != nil {
		return errors.New("invalid UUID format")
	}

	err = s.repo.Delete(parsedUUID)
	if err != nil {
		return errors.New("Failed to delete user")
	}

	return nil
}
