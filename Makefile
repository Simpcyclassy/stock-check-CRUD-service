DOCKER_NAME=stock-check
DOCKER_TAG?=$(shell git rev-parse --short=7 HEAD)
DB_CONTAINER_NAME=postgres-$(DOCKER_TAG)
DB_IMAGE=postgres:10-alpine
DOCKER_BUILD_CONTEXT=./

db: clean start-db migrate

start-db:
	docker run --rm -d \
		--name=$(DB_CONTAINER_NAME) \
		-e POSTGRES_PASSWORD=dev \
		-p 5430:5432 \
		$(DB_IMAGE)

	# Wait for postgres server to start & be ready (accept incoming connections)
	while ! docker exec $(DB_CONTAINER_NAME) pg_isready -h localhost -d postgres -U postgres; do sleep 2; done

ifeq "" "$(shell command -v sql-migrate 2> /dev/null)"
SQL_MIGRATE=docker run --rm --net=container:$(DB_CONTAINER_NAME) $(DOCKER_NAME) sql-migrate
else
SQL_MIGRATE=sql-migrate
endif
# If sql-migrate is available locally, use that - otherwise use the application image
# NOTE: If you're doing local development and don't have sql-migrate installed, you'll
# have to run `make docker-build` before this

migrate:
ifneq "$(SQL_MIGRATE)" "sql-migrate"
	@echo -e "\n\n=> Running migrations in docker, since you don't have 'sql-migrate' installed."
	@echo "=> If this fails, try running 'make docker-build' before running again."
	@echo -e "=> Or, install github.com/rubenv/sql-migrate to make these migrations run faster locally.\n\n"
endif

	$(SQL_MIGRATE) up -config=./database/migrations/user.dbconfig.yml --env local

docker-build:
	docker build \
		-f $(DOCKER_BUILD_CONTEXT)/Dockerfile \
		-t $(DOCKER_NAME) \
		$(DOCKER_BUILD_CONTEXT)

clean:
	docker rm -f $(DB_CONTAINER_NAME) || true
