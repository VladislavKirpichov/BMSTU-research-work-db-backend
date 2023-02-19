build:
	docker-compose build --parallel

up:
	docker-compose up
	migrate -path ./shema -database 'postgres://masha:12345v!@localhost:5432/postgres?sslmode=disable' up

down:
	docker-compose down

start:
	make build
	make up

restart:
	make down
	make start

clean:
	docker builder prune -f
