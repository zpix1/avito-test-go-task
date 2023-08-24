build:
	docker build -t slugs .

run:
	docker run --name slugs -p 8080:8080 -it slugs

remove:
	docker rm -f slugs
	docker rmi slugs

migrate:
	migrate -path ./migrations/postgres/schemas -database postgres://postgres:secret@localhost:5436/slugs?sslmode=disable up

swagger:
	swag init --parseDependency  --parseInternal -g cmd/main.go

postgres:
	docker run --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=slugs -d -v slugs:/var/lib/postgresql/data -p 5436:5432 postgres