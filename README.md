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

`[GET] /books` - get all books

`[POST] /books` - create new book

`[GET] /books/{id}` - get book by id

`[PUT] /books/{id}` - create update book by id

`[DELETE] /books/{id}` - delete book by id

### POST request body example:
```json
{
"title":"book",
"author":"booker",
"rating": 5
}
```


### GET (no books in database):

<img width="764" alt="Снимок экрана 2022-08-22 в 17 32 23" src="https://user-images.githubusercontent.com/88317896/185901248-a6345b63-8bf1-4c89-8792-120b782c5562.png">


### POST (add new book):

<img width="766" alt="Снимок экрана 2022-08-22 в 17 33 12" src="https://user-images.githubusercontent.com/88317896/185901306-897beb11-bc4e-4500-abf1-80ea984139a8.png">


### GET (check new book):
<img width="767" alt="Снимок экрана 2022-08-22 в 17 33 34" src="https://user-images.githubusercontent.com/88317896/185901366-dc77ade4-5e0b-4ea5-b859-c77448911b3b.png">






