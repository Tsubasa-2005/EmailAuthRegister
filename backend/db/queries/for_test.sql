-- users
-- name: TestGetUsers :many
select * from users where id = $1 and deleted_at is null;

-- name: TestGetUsersByEmail :one
select * from users where email = $1 and deleted_at is null;

-- name: TestExistsUserByEmail :one
select exists(select 1 from users where email = $1 and deleted_at is null);

-- name: TestCreateUser :one
insert into users (email, password, name, created_at, updated_at) values ($1, $2, $3, $4, $5) returning *;

-- name: TestUpdateUser :one
update users set email = $1, password = $2, name = $3, updated_at = $4 where id = $5 returning *;

-- name: TestUpdateUserByEmail :one
update users set password = $1, name = $2, updated_at = $3 where email = $4 returning *;

-- name: TestDeleteUser :exec
delete from users where id = $1;

-- email_verification_tokens
-- name: TestGetEmailVerificationToken :one
select * from email_verification_tokens where token = $1;

-- name: TestGetEmailVerificationTokenByEmail :one
select * from email_verification_tokens where email = $1;

-- name: TestCreateEmailVerificationToken :one
insert into email_verification_tokens (token, email, expires_at, created_at) values ($1, $2, $3, $4) returning *;

-- name: TestDeleteEmailVerificationToken :exec
delete from email_verification_tokens where token = $1;
