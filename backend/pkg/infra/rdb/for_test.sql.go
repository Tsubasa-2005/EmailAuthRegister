// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: for_test.sql

package rdb

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const testCreateEmailVerificationToken = `-- name: TestCreateEmailVerificationToken :one
insert into email_verification_tokens (token, email, expires_at, created_at) values ($1, $2, $3, $4) returning token, email, expires_at, created_at
`

type TestCreateEmailVerificationTokenParams struct {
	Token     pgtype.UUID
	Email     string
	ExpiresAt pgtype.Timestamptz
	CreatedAt pgtype.Timestamptz
}

func (q *Queries) TestCreateEmailVerificationToken(ctx context.Context, arg TestCreateEmailVerificationTokenParams) (EmailVerificationToken, error) {
	row := q.db.QueryRow(ctx, testCreateEmailVerificationToken,
		arg.Token,
		arg.Email,
		arg.ExpiresAt,
		arg.CreatedAt,
	)
	var i EmailVerificationToken
	err := row.Scan(
		&i.Token,
		&i.Email,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const testCreateUser = `-- name: TestCreateUser :one
insert into users (email, password, name, created_at, updated_at) values ($1, $2, $3, $4, $5) returning id, email, password, name, created_at, updated_at, deleted_at
`

type TestCreateUserParams struct {
	Email     string
	Password  string
	Name      string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

func (q *Queries) TestCreateUser(ctx context.Context, arg TestCreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, testCreateUser,
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

const testDeleteEmailVerificationToken = `-- name: TestDeleteEmailVerificationToken :exec
delete from email_verification_tokens where token = $1
`

func (q *Queries) TestDeleteEmailVerificationToken(ctx context.Context, token pgtype.UUID) error {
	_, err := q.db.Exec(ctx, testDeleteEmailVerificationToken, token)
	return err
}

const testDeleteUser = `-- name: TestDeleteUser :exec
delete from users where id = $1
`

func (q *Queries) TestDeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, testDeleteUser, id)
	return err
}

const testExistsEmailVerificationTokenByEmail = `-- name: TestExistsEmailVerificationTokenByEmail :one
select exists(select 1 from email_verification_tokens where email = $1)
`

func (q *Queries) TestExistsEmailVerificationTokenByEmail(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRow(ctx, testExistsEmailVerificationTokenByEmail, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const testExistsUserByEmail = `-- name: TestExistsUserByEmail :one
select exists(select 1 from users where email = $1 and deleted_at is null)
`

func (q *Queries) TestExistsUserByEmail(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRow(ctx, testExistsUserByEmail, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const testGetEmailVerificationToken = `-- name: TestGetEmailVerificationToken :one
select token, email, expires_at, created_at from email_verification_tokens where token = $1
`

// email_verification_tokens
func (q *Queries) TestGetEmailVerificationToken(ctx context.Context, token pgtype.UUID) (EmailVerificationToken, error) {
	row := q.db.QueryRow(ctx, testGetEmailVerificationToken, token)
	var i EmailVerificationToken
	err := row.Scan(
		&i.Token,
		&i.Email,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const testGetEmailVerificationTokenByEmail = `-- name: TestGetEmailVerificationTokenByEmail :one
select token, email, expires_at, created_at from email_verification_tokens where email = $1
`

func (q *Queries) TestGetEmailVerificationTokenByEmail(ctx context.Context, email string) (EmailVerificationToken, error) {
	row := q.db.QueryRow(ctx, testGetEmailVerificationTokenByEmail, email)
	var i EmailVerificationToken
	err := row.Scan(
		&i.Token,
		&i.Email,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const testGetUsers = `-- name: TestGetUsers :many
select id, email, password, name, created_at, updated_at, deleted_at from users where id = $1 and deleted_at is null
`

// users
func (q *Queries) TestGetUsers(ctx context.Context, id int64) ([]User, error) {
	rows, err := q.db.Query(ctx, testGetUsers, id)
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

const testGetUsersByEmail = `-- name: TestGetUsersByEmail :one
select id, email, password, name, created_at, updated_at, deleted_at from users where email = $1 and deleted_at is null
`

func (q *Queries) TestGetUsersByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, testGetUsersByEmail, email)
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

const testUpdateUser = `-- name: TestUpdateUser :one
update users set email = $1, password = $2, name = $3, updated_at = $4 where id = $5 returning id, email, password, name, created_at, updated_at, deleted_at
`

type TestUpdateUserParams struct {
	Email     string
	Password  string
	Name      string
	UpdatedAt pgtype.Timestamptz
	ID        int64
}

func (q *Queries) TestUpdateUser(ctx context.Context, arg TestUpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, testUpdateUser,
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

const testUpdateUserByEmail = `-- name: TestUpdateUserByEmail :one
update users set password = $1, name = $2, updated_at = $3 where email = $4 returning id, email, password, name, created_at, updated_at, deleted_at
`

type TestUpdateUserByEmailParams struct {
	Password  string
	Name      string
	UpdatedAt pgtype.Timestamptz
	Email     string
}

func (q *Queries) TestUpdateUserByEmail(ctx context.Context, arg TestUpdateUserByEmailParams) (User, error) {
	row := q.db.QueryRow(ctx, testUpdateUserByEmail,
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
