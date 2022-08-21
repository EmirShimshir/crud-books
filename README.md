# CRUD app for books management

## Tools:
+ go 1.18
+ postgres

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

`[GET] /books/` - get all books

`[POST] /books/` - create new book

`[GET] /books/id/{id}` - get book by id

`[PUT] /books/id/{id}` - create update book by id

`[DELETE] /books/id/{id}` - delete book by id

### POST request body example:
```json
{
"title":"book",
"author":"booker",
"rating": 5
}
```


### GET (no books in database):

<img width="846" alt="Снимок экрана 2022-08-16 в 21 27 21" src="https://user-images.githubusercontent.com/88317896/184955267-177e36df-5cf6-4d8e-94e7-4076224be120.png">

### POST (add new book):

<img width="848" alt="Снимок экрана 2022-08-16 в 21 27 57" src="https://user-images.githubusercontent.com/88317896/184955399-8ac214d6-0e05-4d30-bf9f-968e89ec68e5.png">

### GET (check new book):
<img width="853" alt="Снимок экрана 2022-08-16 в 21 47 23" src="https://user-images.githubusercontent.com/88317896/184955741-df5a8ed2-cc3c-4685-bb3f-8c8900281f32.png">





