package repository

import (
	avito "avito-backend-task-2023"
	"github.com/jmoiron/sqlx"
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
	GetUserSegmentsLogs() ([]avito.UsersSegmentsLogs, error)
}

type Repository struct {
	User
	Segment
	UsersSegments
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:          NewUserPostgres(db),
		Segment:       NewSegmentPostgres(db),
		UsersSegments: NewUsersSegmentsPostgres(db),
	}
}
