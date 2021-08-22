-- +migrate Up

CREATE TABLE stock (
    id                  UUID PRIMARY KEY,
    amount_available	INT DEFAULT 0,
	amount_sold			INT DEFAULT 0,
	description			CHAR(256),
	location			CHAR(256),
	name				CHAR(256) NOT NULL,
	price				FLOAT DEFAULT 0.0
);

-- +migrate Down
DROP TABLE stock;