package repository

import (
	avito "avito-backend-task-2023"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
)

type UsersSegmentsPostgres struct {
	db *sqlx.DB
}

func NewUsersSegmentsPostgres(db *sqlx.DB) *UsersSegmentsPostgres {
	return &UsersSegmentsPostgres{db: db}
}

func (r *UsersSegmentsPostgres) GetUserSegments(userId string) ([]avito.Segment, error) {
	var segments []avito.Segment

	query := fmt.Sprintf("SELECT s.id, s.name FROM %s us INNER JOIN %s s on us.segment_id = s.id WHERE us.user_id = $1",
		usersSegmentsTable, segmentsTable)
	err := r.db.Select(&segments, query, userId)

	return segments, err
}

func (r *UsersSegmentsPostgres) UpdateUserSegments(a avito.AlteredUserSegments) error {
	var segmentsIdToAdd []string
	var segmentsIdToDelete []string

	if len(a.Add) != 0 {
		query, args, err := sqlx.In(fmt.Sprintf("SELECT id FROM %s WHERE name IN (?)", segmentsTable), a.Add)
		if err != nil {
			log.Fatal(err)
		}
		err = r.db.Select(&segmentsIdToAdd, r.db.Rebind(query), args...)
	}

	if len(a.Delete) != 0 {
		query, args, err := sqlx.In(fmt.Sprintf("SELECT id FROM %s WHERE name IN (?)", segmentsTable), a.Delete)
		if err != nil {
			log.Fatal(err)
		}
		err = r.db.Select(&segmentsIdToDelete, r.db.Rebind(query), args...)
	}

	//tx, err := r.db.Begin()
	//if err != nil {
	//	return err
	//}

	if len(a.Delete) != 0 {
		query, args, err := sqlx.In(fmt.Sprintf("DELETE FROM %s WHERE user_id='%s' AND segment_id IN (?)", usersSegmentsTable, a.Id), segmentsIdToDelete)
		_, err = r.db.Exec(r.db.Rebind(query), args...)
		if err != nil {
			//tx.Rollback()
			return err
		}
	}

	if len(a.Add) != 0 {
		vals1 := []interface{}{}
		query := fmt.Sprintf("INSERT INTO %s (id, user_id, segment_id) VALUES", usersSegmentsTable)

		for _, v := range segmentsIdToAdd {
			query += "(?, ?, ?),"
			vals1 = append(vals1, uuid.New().String(), a.Id, v)
		}
		query = strings.TrimSuffix(query, ",")

		query += " ON CONFLICT (user_id, segment_id) DO NOTHING"

		_, err := r.db.Exec(r.db.Rebind(query), vals1...)

		if err != nil {
			//tx.Rollback()
			return err
		}
	}

	return nil
}
