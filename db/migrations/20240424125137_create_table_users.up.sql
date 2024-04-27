create table if not exists "user"
(
    id         text not null
    constraint user_pk
    primary key,
    balance    text,
    role       text,
    created_at timestamp,
    updated_at timestamp
);
