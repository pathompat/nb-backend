package service

import (
	"notebook-backend/handler/dto"
	"notebook-backend/helper"
	"notebook-backend/repository"
	"notebook-backend/repository/model"
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	GetAllUsers() ([]dto.UserResponse, error)
	GetUserByID(userID string) (dto.UserResponse, error)
	GetInfoUser(userID string) (dto.UserResponse, error)
	CreateUser(input dto.CreateUser) (*dto.UserResponse, error)
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
		return nil, err
	}

	userMap := []dto.UserResponse{}
	for _, user := range users {
		userMap = append(userMap, dto.UserResponse{
			UserID:    user.UserID,
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

func (s *userService) GetUserByID(userID string) (dto.UserResponse, error) {
	parsedUUID, err := uuid.Parse(userID)
	if err != nil {
		return dto.UserResponse{}, err
	}

	user, err := s.repo.FindByID(parsedUUID)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		UserID:    user.UserID,
		TierID:    user.TierID,
		Username:  user.Username,
		StoreName: user.StoreName,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *userService) GetInfoUser(userID string) (dto.UserResponse, error) {
	parsedUUID, err := uuid.Parse(userID)
	if err != nil {
		return dto.UserResponse{}, err
	}

	user, err := s.repo.FindByID(parsedUUID)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		UserID:    user.UserID,
		TierID:    user.TierID,
		Username:  user.Username,
		StoreName: user.StoreName,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *userService) CreateUser(input dto.CreateUser) (*dto.UserResponse, error) {
	hashPassword, err := helper.HashPassword(input.Password)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return &dto.UserResponse{
		UserID:    createdUser.UserID,
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
		return dto.UserResponse{}, err
	}

	user, err := s.repo.FindByID(parsedUUID)
	if err != nil {
		return dto.UserResponse{}, err
	}

	if input.Password != nil && *input.Password != "" {
		hashPassword, err := helper.HashPassword(*input.Password)
		if err != nil {
			return dto.UserResponse{}, helper.ErrHashPassword
		}
		user.Password = hashPassword
	}

	user.Username = input.Username
	user.TierID = input.TierID
	user.StoreName = input.StoreName
	user.UpdatedAt = time.Now()

	updateUser, err := s.repo.Update(user)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		UserID:    updateUser.UserID,
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
		return err
	}

	err = s.repo.Delete(parsedUUID)
	if err != nil {
		return err
	}

	return nil
}
