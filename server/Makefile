postgresinit:
	docker run --name postgresCardgame -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:latest

postgres:
	docker exec -it postgresCardgame psql

createdb:
	docker exec -it postgresCardgame createdb --username=root --owner=root cardgame_db

dropdb:
	docker exec -tr postgresCardgame dropdb cardgame_db


.PHONY: postgresinit postgres createdb dropdb