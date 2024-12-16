// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package rdb

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
insert into users (email, password, name, created_at, updated_at) values ($1, $2, $3, $4, $5) returning id, email, password, name, created_at, updated_at, deleted_at
`

type CreateUserParams struct {
	Email     string
	Password  string
	Name      string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Email,
		arg.Password,
		arg.Name,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const existsUserByEmail = `-- name: ExistsUserByEmail :one
select exists(select 1 from users where email = $1 and deleted_at is null)
`

func (q *Queries) ExistsUserByEmail(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRow(ctx, existsUserByEmail, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getUsers = `-- name: GetUsers :many
select id, email, password, name, created_at, updated_at, deleted_at from users where id = $1 and deleted_at is null
`

func (q *Queries) GetUsers(ctx context.Context, id int64) ([]User, error) {
	rows, err := q.db.Query(ctx, getUsers, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Password,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUsersByEmail = `-- name: GetUsersByEmail :one
select id, email, password, name, created_at, updated_at, deleted_at from users where email = $1 and deleted_at is null
`

func (q *Queries) GetUsersByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUsersByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const softDeleteUsers = `-- name: SoftDeleteUsers :exec
update users set deleted_at = $1 where id = $2
`

type SoftDeleteUsersParams struct {
	DeletedAt pgtype.Timestamptz
	ID        int64
}

func (q *Queries) SoftDeleteUsers(ctx context.Context, arg SoftDeleteUsersParams) error {
	_, err := q.db.Exec(ctx, softDeleteUsers, arg.DeletedAt, arg.ID)
	return err
}

const updateUser = `-- name: UpdateUser :one
update users set email = $1, password = $2, name = $3, updated_at = $4 where id = $5 returning id, email, password, name, created_at, updated_at, deleted_at
`

type UpdateUserParams struct {
	Email     string
	Password  string
	Name      string
	UpdatedAt pgtype.Timestamptz
	ID        int64
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Email,
		arg.Password,
		arg.Name,
		arg.UpdatedAt,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const updateUserByEmail = `-- name: UpdateUserByEmail :one
update users set password = $1, name = $2, updated_at = $3 where email = $4 returning id, email, password, name, created_at, updated_at, deleted_at
`

type UpdateUserByEmailParams struct {
	Password  string
	Name      string
	UpdatedAt pgtype.Timestamptz
	Email     string
}

func (q *Queries) UpdateUserByEmail(ctx context.Context, arg UpdateUserByEmailParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUserByEmail,
		arg.Password,
		arg.Name,
		arg.UpdatedAt,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
