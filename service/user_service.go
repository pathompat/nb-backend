package service

import (
	"errors"
	"notebook-backend/controller/dto"
	"notebook-backend/repository"
	"notebook-backend/repository/model"
	"time"
)

type UserService interface {
	GetAllUsers() ([]dto.User, error)
	CreateUser(input dto.CreateUserDTO) (dto.User, error)
	UpdateUser(userID string, input dto.UpdateUserDTO) (dto.User, error)
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
			ID:        int(user.ID),
			Username:  user.Username,
			StoreName: user.StoreName,
			TierID:    user.TierID,
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
		ID:        int(createdUser.ID),
		TierID:    createdUser.TierID,
		Username:  createdUser.Username,
		StoreName: createdUser.StoreName,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
	}, nil
}

func (s *userService) UpdateUser(userID string, input dto.UpdateUserDTO) (dto.User, error) {
	user, err := s.repo.FindByID(userID)
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
		ID:        int(updateUser.ID),
		TierID:    updateUser.TierID,
		Username:  updateUser.Username,
		StoreName: updateUser.StoreName,
		CreatedAt: updateUser.CreatedAt,
		UpdatedAt: updateUser.UpdatedAt,
	}, nil
}
