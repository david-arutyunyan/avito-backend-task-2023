package service

import (
	avito "avito-backend-task-2023"
	"avito-backend-task-2023/pkg/repository"
	"github.com/gocarina/gocsv"
	"net/http"
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

func (s *UsersSegmentsService) GetUserSegmentsLogs(rw http.ResponseWriter, date avito.CustomDate) error {
	logs, err := s.repo.GetUserSegmentsLogs(date)

	err = gocsv.Marshal(logs, rw)
	if err != nil {
		return err
	}

	return err
}
