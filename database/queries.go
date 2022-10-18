package database

import (
	"context"
	"user_assignment/model"
)

// Create
const createUser = `-- name: CreateUser :one
INSERT INTO users (name)
VALUES ($1)
RETURNING id, name
`

func (q *Queries) CreateUser(ctx context.Context, name string) (model.User, error) {
	row := q.db.QueryRowContext(ctx, createUser, name)
	var i model.User
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

// Delete
const deleteUser = `-- name: DeleteUser :one
DELETE
FROM users
WHERE id = $1
RETURNING *
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) (model.User, error) {
	//_, err := q.db.ExecContext(ctx, deleteUser, id)
	//return err
	row := q.db.QueryRowContext(ctx, deleteUser, id)
	var i model.User
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

// Get
const getUser = `-- name: GetUser :one
SELECT id, name
FROM users
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (model.User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i model.User
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

// List
const getUserList = `-- name: GetUserList :many
SELECT id, name
FROM users
ORDER BY id
`

func (q *Queries) GetUserList(ctx context.Context) ([]model.User, error) {
	rows, err := q.db.QueryContext(ctx, getUserList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.User
	for rows.Next() {
		var i model.User
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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

// Update
const updateUser = `-- name: UpdateUser :one
UPDATE users
SET name = $2
WHERE id = $1
RETURNING id, name
`

type UpdateUserParams struct {
	ID   int32  `json:"user_id,numeric"`
	Name string `json:"name" binding:"required,max=20,alpha"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (model.User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.ID, arg.Name)
	var i model.User
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

// Truncate
const truncateUser = `-- name: TruncateUser :exec
TRUNCATE users
`

func (q *Queries) TruncateUsers(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, truncateUser)
	return err
}
