package service

import (
	"errors"
	"notebook-backend/controller/dto"
	"notebook-backend/repository"
	"notebook-backend/repository/model"
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	GetAllUsers() ([]dto.User, error)
	CreateUser(input dto.CreateUserDTO) (dto.User, error)
	UpdateUser(userID string, input dto.UpdateUserDTO) (dto.User, error)
	DeleteUser(userID string) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers() ([]dto.User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, errors.New("database error")
	}

	userMap := []dto.User{}
	for _, user := range users {
		userMap = append(userMap, dto.User{
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

func (s *userService) CreateUser(input dto.CreateUserDTO) (dto.User, error) {
	newUser := model.User{
		Username:  input.Username,
		TierID:    input.TierID,
		StoreName: input.StoreName,
		Password:  input.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdUser, err := s.repo.Create(newUser)
	if err != nil {
		return dto.User{}, errors.New("Failed to create user")
	}

	return dto.User{
		UserId:    createdUser.UserId,
		TierID:    createdUser.TierID,
		Username:  createdUser.Username,
		StoreName: createdUser.StoreName,
		Role:      createdUser.Role,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
	}, nil
}

func (s *userService) UpdateUser(userID string, input dto.UpdateUserDTO) (dto.User, error) {
	parsedUUID, err := uuid.Parse(userID)
	if err != nil {
		return dto.User{}, errors.New("invalid UUID format")
	}

	user, err := s.repo.FindByID(parsedUUID)
	if err != nil {
		return dto.User{}, errors.New("User not found")
	}

	user.Username = input.Username
	user.Password = input.Password
	user.TierID = input.TierID
	user.StoreName = input.StoreName
	user.UpdatedAt = time.Now()

	updateUser, err := s.repo.Update(user)
	if err != nil {
		return dto.User{}, errors.New("Failed to update user")
	}

	return dto.User{
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
