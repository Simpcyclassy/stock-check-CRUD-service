# stock-check-CRUD-service
A platform to creates and manages stocks

## Run Server

Run this command in the terminal

```
go run main.go
```

- [x] Add Go modules
- [x] Define the models
- [x] Create a server
- [x] Create 4 handlers and a print naming what each handler does eg (create, delete, update, read)
- [ ] Set up middleware / validators
- [ ] Need a database, postgres?
- [ ] Database package ie package to use in accessing database
- [ ] Run database migrations
- [ ] Add helpers like linting, and configurations for variables
- [ ] Set up Docker
- [ ] Set up Grafana and prometheus
- [ ] Set up logger eg zerolog


## Database

Create the database locally:

```sh
make db
# you may also need to run `make docker-build` the first time
```

If the port is already allocated you will need to run:

```sh
lsof -i :5432
# to find out what is running at the port 5432, then
kill <pid-number>
# to kill the running process blocking port 5432
# sometimes you may have to do this `sudo`
```

You can inspect the database by connecting manually using the [psql tool](https://www.postgresql.org/docs/9.2/app-psql.html), [installation instructions](https://blog.timescale.com/tutorials/how-to-install-psql-on-mac-ubuntu-debian-windows/). Use `psql -U postgres -h localhost` to connect (after the database is running locally), then enter the local-dev password (which is dev). Inside psql, the command `\dt` will show you a list of tables, and `\d <table_name>` will describe a table.

Example:

```psql
postgres-# \d stock
                         Table "public.stock"
      Column      |       Type       | Collation | Nullable | Default 
------------------+------------------+-----------+----------+---------
 id               | uuid             |           | not null | 
 amount_available | integer          |           |          | 0
 amount_sold      | integer          |           |          | 0
 description      | character(256)     |           |          | 
 location         | character(256)     |           |          | 
 name             | character(256)     |           | not null | 
 price            | double precision |           |          | 
Indexes:
    "stock_pkey" PRIMARY KEY, btree (id)
```

To add data into the database manually you can use the following SQL command inside the psql tool

```psql
insert into stock(id, amount_available, amount_sold, available, description, location, name, price) VALUES('956c7886-4243-4f82-8d30-7b92d26a8fce', 65, 5, 'comfortable skaters', 'berlin);
```