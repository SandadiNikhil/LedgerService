#  Makefile for managing PostgreSQL, migrations, and code generation
postgres:
	docker run --name postgres18 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:18-alpine

#  To create the database
createdb:
	docker exec -it postgres18 createdb --username=root --owner=root simple_bank

#  To drop the database
dropdb:
	docker exec -it postgres18 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

#  To rollback the last migration
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

#  Generate SQL code:
sqlc:
	sqlc generate

#  Run tests (to verify everything is working):
test :
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test stop-postgres logs-postgres



#  To stop and remove the postgres container
stop-postgres:
	docker stop postgres18 && docker rm postgres18

#  To view logs of the postgres container
logs-postgres:
	docker logs -f postgres18


#  Commands to Start the Project
# 	make postgres
# 	make createdb
# 	make migrateup
# 	make sqlc
# 	make test

# 	Commands to Stop the Project
# 	make dropdb
# 	make migratedown
# 	make stop-postgres