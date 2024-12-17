-- name: GetUsers :many
select * from users where deleted_at is null limit $1 offset $2;

-- name: GetUserByEmail :one
select * from users where email = $1 and deleted_at is null;

-- name: ExistsUserByEmail :one
select exists(select 1 from users where email = $1 and deleted_at is null);

-- name: CountActiveUsers :one
select count(*) from users where deleted_at is null;

-- name: CreateUser :one
insert into users (email, password, name, created_at, updated_at) values ($1, $2, $3, $4, $5) returning *;

-- name: UpdateUser :one
update users set email = $1, password = $2, name = $3, updated_at = $4 where id = $5 returning *;

-- name: UpdateUserByEmail :one
update users set password = $1, name = $2, updated_at = $3 where email = $4 returning *;

-- name: SoftDeleteUsers :exec
update users set deleted_at = $1 where id = $2;