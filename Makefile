include .env
MIGRATIONS_FOLDER = ./database/migrations

migrate.create:
	@if [ "$(name)" = "" ]; then \
		echo "Error: Please provide a migration name using 'name=your_migration_name'"; \
		exit 1; \
	fi
	migrate create -ext sql -dir $(MIGRATIONS_FOLDER) $(name)
migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down

migrate.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(version)