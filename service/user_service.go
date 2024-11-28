package service

import (
	"errors"
	"notebook-backend/controller/dto"
	"notebook-backend/repository"
)

type UserService interface {
	GetAllUsers() ([]dto.User, error)
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
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return userMap, nil
}
