include ./.env

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_DATA)?sslmode=disable
MIGRATION_PATH=db/migrations

migrate-up:
	migrate -database "$(DB_URL)" -path $(MIGRATION_PATH) up

migrate-down:
	migrate -database "$(DB_URL)" -path $(MIGRATION_PATH) down 1

migrate-version:
	migrate -database "$(DB_URL)" -path $(MIGRATION_PATH) version

migrate-force:
	migrate -database "$(DB_URL)" -path $(MIGRATION_PATH) force

migrate-create:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq create_$(nt)_table

seed:
	psql "$(DB_URL)" -f db/seeding/001_users.sql

pg:
	pg_dump "$(DB_URL)" --data-only --column-inserts --table=users --file=./db/seeding/001_users.sql