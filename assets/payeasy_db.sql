CREATE DATABASE IF NOT EXISTS 'payeasy';
CREATE EXTENSION IF NOT EXISTS 'uuid-ossp';
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TYPE role_type AS ENUM ('customer', 'merchant');
CREATE TYPE status_type AS ENUM ('credit', 'debit');

CREATE TABLE users(
	id uuid_generate_v4(),
	name VARCHAR(50) NOT NULL,
	email VARCHAR(50) NOT NULL,
	password VARCHAR(50) NOT NULL,
	number VARCHAR NOT NULL,
	address VARCHAR(50) NOT NULL,
	role role_type NOT NULL,
	balance BIGINT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE merchants(
	id uuid_generate_v4() NOT NULL,
	name_merchants VARCHAR(50) NOT NULL,
	id_users uuid NOT NULL,
	balance BIGINT,
	FOREIGN KEY (id_users) REFERENCES users(id),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE trx_history(
	id uuid_generate_v4() NOT NULL,
	id_users uuid NOT NULL,
	id_merchants uuid NOT NULL,
	status_payment status_type,
	total_amount BIGINT,
	FOREIGN KEY (id_users) REFERENCES users(id),
	FOREIGN KEY (id_merchants) REFERENCES merchants(id),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);