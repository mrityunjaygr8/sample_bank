ifneq (,$(wildcard ./.env))
    include .env
    export
endif

cmd-exists-%:
	@hash $(*) > /dev/null 2>&1 || \
		(echo "ERROR: '$(*)' must be installed and available on your PATH."; exit 1)
dbstart: cmd-exists-docker
	docker run --name postgres13 -e POSTGRES_PASSWORD="${DB_PASS}" -e POSTGRES_USER="${DB_USER}" -d  -p "${DB_PORT}":5432 postgres:13-alpine

dbstop:
	docker stop postgres13
	docker rm postgres13

dbcreate:
	docker exec -it postgres13 createdb --username="${DB_USER}" --owner="${DB_USER}" "${DB_NAME}"

dbdrop:
	docker exec -it postgres13 dropdb "${DB_NAME}"

migrateup: cmd-exists-migrate
	migrate -path db/migration -database "postgresql://"${DB_USER}":"${DB_PASS}"@"${DB_HOST}":"${DB_PORT}"/"${DB_NAME}"?sslmode=disable" -verbose up


migratedown: cmd-exists-migrate
	migrate -path db/migration -database "postgresql://"${DB_USER}":"${DB_PASS}"@"${DB_HOST}":"${DB_PORT}"/"${DB_NAME}"?sslmode=disable" -verbose down

sqlc: cmd-exists-sqlc
	sqlc generate

.PHONY: dbstart dbstop dbcreate dbdrop migrateup migrateup sqlc
