create table users (
	id bigserial primary key,
	email varchar(255) unique not null,
	password varchar(64) not null,
	name varchar(255) not null,
	created_at timestamptz not null,
	updated_at timestamptz not null,
	deleted_at timestamptz
);

create table email_verification_tokens (
	token uuid primary key,
	email varchar(255) not null,
	expires_at timestamptz not null,
	created_at timestamptz not null
);