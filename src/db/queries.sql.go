// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: queries.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todo WHERE id = $1
`

func (q *Queries) DeleteTodo(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTodo, id)
	return err
}

const insertProfile = `-- name: InsertProfile :one
INSERT INTO profile (user_id, username) values ($1, $2) RETURNING id, username, user_id, todos
`

type InsertProfileParams struct {
	UserID   uuid.UUID
	Username string
}

func (q *Queries) InsertProfile(ctx context.Context, arg InsertProfileParams) (Profile, error) {
	row := q.db.QueryRowContext(ctx, insertProfile, arg.UserID, arg.Username)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.UserID,
		&i.Todos,
	)
	return i, err
}

const insertTodo = `-- name: InsertTodo :one
INSERT INTO todo (body) values ($1) RETURNING id, created_at, body
`

func (q *Queries) InsertTodo(ctx context.Context, body string) (Todo, error) {
	row := q.db.QueryRowContext(ctx, insertTodo, body)
	var i Todo
	err := row.Scan(&i.ID, &i.CreatedAt, &i.Body)
	return i, err
}

const selectEmailAlreadyExists = `-- name: SelectEmailAlreadyExists :one
SELECT email FROM auth.users WHERE auth.users.email = $1
`

func (q *Queries) SelectEmailAlreadyExists(ctx context.Context, email sql.NullString) (sql.NullString, error) {
	row := q.db.QueryRowContext(ctx, selectEmailAlreadyExists, email)
	err := row.Scan(&email)
	return email, err
}

const selectProfileByAuthUserId = `-- name: SelectProfileByAuthUserId :one
SELECT id FROM profile WHERE profile.user_id = $1
`

func (q *Queries) SelectProfileByAuthUserId(ctx context.Context, userID uuid.UUID) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectProfileByAuthUserId, userID)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectProfileById = `-- name: SelectProfileById :one
SELECT id, username, user_id, todos FROM profile WHERE profile.id = $1
`

func (q *Queries) SelectProfileById(ctx context.Context, id int64) (Profile, error) {
	row := q.db.QueryRowContext(ctx, selectProfileById, id)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.UserID,
		&i.Todos,
	)
	return i, err
}

const selectProfileByUsername = `-- name: SelectProfileByUsername :one
SELECT id, username, user_id, todos FROM profile WHERE profile.username = $1
`

func (q *Queries) SelectProfileByUsername(ctx context.Context, username string) (Profile, error) {
	row := q.db.QueryRowContext(ctx, selectProfileByUsername, username)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.UserID,
		&i.Todos,
	)
	return i, err
}

const selectTodos = `-- name: SelectTodos :many
SELECT id, created_at, body FROM todo LIMIT $1
`

func (q *Queries) SelectTodos(ctx context.Context, limit int64) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, selectTodos, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(&i.ID, &i.CreatedAt, &i.Body); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
