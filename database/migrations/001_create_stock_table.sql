-- +migrate Up

CREATE TABLE stock (
    id                  UUID PRIMARY KEY,
    amount_available	INT DEFAULT 0,
	amount_sold			INT DEFAULT 0,
	available			BOOL DEFAULT false,
	description			CHAR,
	location			CHAR,
	name				CHAR NOT NULL,
	price				FLOAT DEFAULT 0.0
);

-- +migrate Down
DROP TABLE stock;