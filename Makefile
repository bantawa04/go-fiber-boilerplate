include .env
MIGRATIONS_FOLDER = ./database/migrations
DB_DSN="postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}"
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

dao:
		@command -v gentool >/dev/null 2>&1 || (echo "Installing gentool..." && go install gorm.io/gen/tools/gentool@latest)
		gentool -db postgres -dsn ${DB_DSN} -fieldNullable -fieldWithIndexTag -fieldWithTypeTag -fieldSignable -onlyModel -outPath "./app/dao" -modelPkgName "dao"

.PHONY: dao migrate.create migrate.down migrate.force migrate.up

