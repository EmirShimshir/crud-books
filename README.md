# CRUD app for books management

## Tools:
+ go 1.19
+ postgres
+ gin
+ swagger
+ cache

## How to start:
+ Install docker-compose and make

+ Run commands:

```
git clone https://github.com/EmirShimshir/crud-books.git
```

```
make run
```

## API example

### This REST API contains the following methods:

`[GET] /books` - get all books

`[POST] /books` - create new book

`[GET] /books/{id}` - get book by id

`[PUT] /books/{id}` - create update book by id

`[DELETE] /books/{id}` - delete book by id
