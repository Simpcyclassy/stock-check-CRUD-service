# stock-check-CRUD-service
A platform to creates and manages stocks

## Run Server

Run this command in the terminal

```
go run main.go
```

[x] Add Go modules
[x] Define the models
[x] Create a server
[x] Create 4 handlers and a print naming what each handler does eg (create, delete, update, read)
[ ] Set up middleware / validators
[ ] Need a database, postgres?
[ ] Database package ie package to use in accessing database
[ ] Run database migrations
[ ] Add helpers like linting, and configurations for variables
[ ] Set up Docker
[ ] Set up Grafana and prometheus
[ ] Set up logger eg zerolog


## Database

Create the database locally:

```sh
make db
# you may also need to run `make docker-build` the first time
```

You can inspect the database by connecting manually using the [psql tool](https://www.postgresql.org/docs/9.2/app-psql.html), [installation instructions](https://blog.timescale.com/tutorials/how-to-install-psql-on-mac-ubuntu-debian-windows/). Use `psql -U postgres -h localhost` to connect (after the database is running locally), then enter the local-dev password (which is dev). Inside psql, the command `\dt` will show you a list of tables, and `\d <table_name>` will describe a table.

Example:

```
postgres-# \d stock
                         Table "public.stock"
      Column      |       Type       | Collation | Nullable | Default 
------------------+------------------+-----------+----------+---------
 id               | uuid             |           | not null | 
 amount_available | integer          |           |          | 0
 amount_sold      | integer          |           |          | 0
 available        | boolean          |           |          | false
 description      | character(1)     |           |          | 
 location         | character(1)     |           |          | 
 name             | character(1)     |           | not null | 
 price            | double precision |           |          | 
Indexes:
    "stock_pkey" PRIMARY KEY, btree (id)
```