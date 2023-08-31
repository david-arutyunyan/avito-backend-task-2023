package repository

import (
	avito "avito-backend-task-2023"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
	"time"
)

type UsersSegmentsPostgres struct {
	db *sqlx.DB
}

func NewUsersSegmentsPostgres(db *sqlx.DB) *UsersSegmentsPostgres {
	return &UsersSegmentsPostgres{db: db}
}

func (r *UsersSegmentsPostgres) GetUserSegmentsLogs() ([]avito.UsersSegmentsLogs, error) {
	var segments []avito.UsersSegmentsLogs

	query := fmt.Sprintf("SELECT id, user_id as userId, segment_name as segmentName, operation, time FROM %s", usersSegmentsLogsTable)
	err := r.db.Select(&segments, query)
	if err != nil {
		log.Fatal(err)
	}

	return segments, err
}

func (r *UsersSegmentsPostgres) GetUserSegments(userId string) ([]avito.Segment, error) {
	var segments []avito.Segment

	query := fmt.Sprintf("SELECT s.id, s.name FROM %s us INNER JOIN %s s on us.segment_id = s.id WHERE us.user_id = $1",
		usersSegmentsTable, segmentsTable)
	err := r.db.Select(&segments, query, userId)

	return segments, err
}

func (r *UsersSegmentsPostgres) UpdateUserSegments(userSegments avito.AlteredUserSegments) error {
	var segmentsIdToAdd []string
	var segmentsIdToDelete []string

	if len(userSegments.Add) != 0 {
		query, args, err := sqlx.In(fmt.Sprintf("SELECT id FROM %s WHERE name IN (?)", segmentsTable), userSegments.Add)
		if err != nil {
			log.Fatal(err)
		}
		err = r.db.Select(&segmentsIdToAdd, r.db.Rebind(query), args...)
	}

	if len(userSegments.Delete) != 0 {
		query, args, err := sqlx.In(fmt.Sprintf("SELECT id FROM %s WHERE name IN (?)", segmentsTable), userSegments.Delete)
		if err != nil {
			log.Fatal(err)
		}
		err = r.db.Select(&segmentsIdToDelete, r.db.Rebind(query), args...)
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if len(userSegments.Delete) != 0 {
		query, args, err := sqlx.In(fmt.Sprintf("DELETE FROM %s WHERE user_id='%s' AND segment_id IN (?)", usersSegmentsTable, userSegments.Id), segmentsIdToDelete)
		_, err = tx.Exec(r.db.Rebind(query), args...)
		if err != nil {
			tx.Rollback()
			return err
		}

		err = updateUserSegmentsLogs(tx, r, userSegments.Id, userSegments.Delete, "DELETE", time.Now())
		if err != nil {
			tx.Rollback()
			return err
		}

	}

	if len(userSegments.Add) != 0 {
		vals := []interface{}{}
		query := fmt.Sprintf("INSERT INTO %s (id, user_id, segment_id) VALUES", usersSegmentsTable)

		for _, v := range segmentsIdToAdd {
			query += "(?, ?, ?),"
			vals = append(vals, uuid.New().String(), userSegments.Id, v)
		}
		query = strings.TrimSuffix(query, ",")

		query += " ON CONFLICT (user_id, segment_id) DO NOTHING"

		_, err := tx.Exec(r.db.Rebind(query), vals...)

		if err != nil {
			tx.Rollback()
			return err
		}

		err = updateUserSegmentsLogs(tx, r, userSegments.Id, userSegments.Add, "ADD", time.Now())
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func updateUserSegmentsLogs(tx *sql.Tx, r *UsersSegmentsPostgres, userId string, segments []string, operation string, time time.Time) error {
	vals := []interface{}{}
	query := fmt.Sprintf("INSERT INTO %s (id, user_id, segment_name, operation, time) VALUES", usersSegmentsLogsTable)

	for _, name := range segments {
		query += "(?, ?, ?, ?, ?),"
		vals = append(vals, uuid.New().String(), userId, name, operation, time)
	}
	query = strings.TrimSuffix(query, ",")

	_, err := tx.Exec(r.db.Rebind(query), vals...)

	if err != nil {
		tx.Rollback()
		return err
	}

	return err
}
