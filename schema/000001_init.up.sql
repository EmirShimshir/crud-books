create TABLE books
(
    id           serial       not null unique,
    title        varchar(255) not null unique,
    author       varchar(255) not null,
    publish_date timestamp    not null default now(),
    rating       int          not null
)