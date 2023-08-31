package service

import (
	avito "avito-backend-task-2023"
	"avito-backend-task-2023/pkg/repository"
)

type UsersSegmentsService struct {
	repo repository.UsersSegments
}

func NewUsersSegmentsService(repo repository.UsersSegments) *UsersSegmentsService {
	return &UsersSegmentsService{repo: repo}
}

func (s *UsersSegmentsService) GetUserSegments(userId string) ([]avito.Segment, error) {
	return s.repo.GetUserSegments(userId)
}

func (s *UsersSegmentsService) UpdateUserSegments(a avito.AlteredUserSegments) error {
	return s.repo.UpdateUserSegments(a)
}
