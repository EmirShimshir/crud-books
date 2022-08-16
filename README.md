# CRUD app for books management

## Tools:
+ go 1.18
+ postgres

## How to start:
Run docker container with postgres:
```
docker run -d --name books -e POSTGRES_PASSWORD=qwerty123 -v {$HOME}/pgdata/:/var/lib/postgresql/data -p 5432:5432 postgres
```

Start container terminal and use psql:
```
docker exec -it books bash
psql -U postgres
```

Create a table for books:
```sql
create table books (
id serial not null unique,
title varchar(255) not null unique,
author varchar(255) not null,
publish_date timestamp not null default now(),
rating int not null
);
```

Run application:
```
go run cmd/main.go
```

## API example