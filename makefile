migrate_from_app:
	migrate -path ./schema -database 'postgres://postgres:qwerty@db:5432/postgres?sslmode=disable' up

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

build:
	docker-compose up --build crud-books

run:
	docker-compose up crud-books

clean:
	docker stop crud-books_db_1 || true
	docker stop stop crud-books_crud-books_1 || true
	docker rm crud-books_db_1 || true
	docker rm crud-books_crud-books_1 || true
	docker rmi crud-books_crud-books || true
	docker rmi postgres || true

clean_with_db:
	make clean
	docker volume rm crud-books_postgres_volume || true



