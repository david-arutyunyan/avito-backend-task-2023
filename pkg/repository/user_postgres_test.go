package repository

//
//import (
//	"errors"
//	"github.com/google/uuid"
//	"github.com/stretchr/testify/assert"
//	sqlmock "github.com/zhashkevych/go-sqlxmock"
//	"testing"
//
//	avito "avito-backend-task-2023"
//)
//
////func (r *UserPostgres) Test_CreateUsera(user avito.User) (string, error) {
////	var id string
////	query := fmt.Sprintf("INSERT INTO %s (id, name, username, password_hash) VALUES ($1, $2, $3, $4) RETURNING id", usersTable)
////
////	row := r.db.QueryRow(query, uuid.New().String(), user.Name, user.Username, user.Password)
////	if err := row.Scan(&id); err != nil {
////		return "", err
////	}
////
////	return id, nil
////}
//
//func Test_CreateUser(t *testing.T) {
//	db, mock, err := sqlmock.Newx()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	r := NewUserPostgres(db)
//
//	k := uuid.New().String()
//
//	type args struct {
//		item avito.User
//	}
//
//	type mockBehavior func(args args)
//
//	tests := []struct {
//		name    string
//		mock    mockBehavior
//		input   args
//		want    string
//		wantErr bool
//	}{
//		{
//			name: "Ok",
//			input: args{
//				item: avito.User{
//					Id:       k,
//					Username: "test_username",
//					Name:     "test_name",
//					Password: "test_password",
//				},
//			},
//			want: k,
//			mock: func(args args) {
//				rows := sqlmock.NewRows([]string{"id"}).AddRow(args.item.Id)
//				mock.ExpectQuery("INSERT INTO users").
//					WithArgs(args.item.Id, args.item.Username, args.item.Name, args.item.Password).WillReturnRows(rows)
//			},
//		},
//		{
//			name: "Empty Fields",
//			input: args{
//				item: avito.User{},
//			},
//			mock: func(args args) {
//				mock.ExpectBegin()
//
//				rows := sqlmock.NewRows([]string{"id"}).AddRow(0).RowError(0, errors.New("insert error"))
//				mock.ExpectQuery("INSERT INTO users").
//					WithArgs(args.item.Id, args.item.Username, args.item.Name, args.item.Password).WillReturnRows(rows)
//
//				mock.ExpectRollback()
//			},
//			wantErr: true,
//		},
//		//{
//		//	name: "Failed 2nd Insert",
//		//	input: args{
//		//		listId: 1,
//		//		item: avito.TodoItem{
//		//			Title:       "title",
//		//			Description: "description",
//		//		},
//		//	},
//		//	mock: func(args args, id int) {
//		//		mock.ExpectBegin()
//		//
//		//		rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
//		//		mock.ExpectQuery("INSERT INTO todo_items").
//		//			WithArgs(args.item.Title, args.item.Description).WillReturnRows(rows)
//		//
//		//		mock.ExpectExec("INSERT INTO lists_items").WithArgs(args.listId, id).
//		//			WillReturnError(errors.New("insert error"))
//		//
//		//		mock.ExpectRollback()
//		//	},
//		//	wantErr: true,
//		//},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.mock(tt.input)
//
//			got, err := r.CreateUser(tt.input.item)
//			if tt.wantErr {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//				assert.Equal(t, tt.want, got)
//			}
//			assert.NoError(t, mock.ExpectationsWereMet())
//		})
//	}
//}
