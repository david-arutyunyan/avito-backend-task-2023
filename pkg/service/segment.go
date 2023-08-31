package service

import (
	avito "avito-backend-task-2023"
	"avito-backend-task-2023/pkg/repository"
)

type SegmentService struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *SegmentService {
	return &SegmentService{repo: repo}
}

func (s *SegmentService) CreateSegment(segment avito.Segment) (string, error) {
	return s.repo.CreateSegment(segment)
}

func (s *SegmentService) DeleteSegment(segment avito.Segment) error {
	return s.repo.DeleteSegment(segment)
}
