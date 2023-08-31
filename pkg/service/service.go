package service

import (
	avito "avito-backend-task-2023"
	"avito-backend-task-2023/pkg/repository"
	"net/http"
)

type User interface {
	CreateUser(user avito.User) (string, error)
	DeleteUser(userId string) error
}

type Segment interface {
	CreateSegment(segment avito.Segment) (string, error)
	DeleteSegment(segment avito.Segment) error
}

type UsersSegments interface {
	GetUserSegments(userId string) ([]avito.Segment, error)
	UpdateUserSegments(a avito.AlteredUserSegments) error
	GetUserSegmentsLogs(rw http.ResponseWriter) error
}

type Service struct {
	User
	Segment
	UsersSegments
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:          NewUserService(repo.User),
		Segment:       NewSegmentService(repo.Segment),
		UsersSegments: NewUsersSegmentsService(repo.UsersSegments),
	}
}
