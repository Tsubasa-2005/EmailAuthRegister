-- name: GetEmailVerificationToken :one
select * from email_verification_tokens where token = $1;

-- name: GetEmailVerificationTokenByEmail :one
select * from email_verification_tokens where email = $1;

-- name: CreateEmailVerificationToken :one
insert into email_verification_tokens (token, email, expires_at, created_at) values ($1, $2, $3, $4) returning *;

-- name: DeleteEmailVerificationToken :exec
delete from email_verification_tokens where token = $1;

-- name: DeleteEmailVerificationTokensByEmail :exec
delete from email_verification_tokens where email = $1;