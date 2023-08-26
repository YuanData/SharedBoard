DB_URL=postgresql://root:helloworld@localhost:5432/sharedboard?sslmode=disable

network:
	docker network create trade-network
postgres:
	docker run --name postgres --network trade-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=helloworld -d postgres:14-alpine

create_db:
	docker exec -it postgres createdb --username=root --owner=root sharedboard
drop_db:
	docker exec -it postgres dropdb sharedboard

migrate_up:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
migrate_down:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

.PHONY: network postgres create_db drop_db migrate_up migrate_down