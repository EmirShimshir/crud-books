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

### GET (no books in database):

<img width="846" alt="Снимок экрана 2022-08-16 в 21 27 21" src="https://user-images.githubusercontent.com/88317896/184955267-177e36df-5cf6-4d8e-94e7-4076224be120.png">

### POST (add new book):

<img width="848" alt="Снимок экрана 2022-08-16 в 21 27 57" src="https://user-images.githubusercontent.com/88317896/184955399-8ac214d6-0e05-4d30-bf9f-968e89ec68e5.png">

### GET (check new book):
<img width="853" alt="Снимок экрана 2022-08-16 в 21 47 23" src="https://user-images.githubusercontent.com/88317896/184955741-df5a8ed2-cc3c-4685-bb3f-8c8900281f32.png">





