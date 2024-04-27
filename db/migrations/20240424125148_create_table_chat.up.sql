create table if not exists chat
(
    id         text not null
    constraint chat_pk
    primary key,
    user_fk    text
    constraint user_fk
    references "user",
    message    text not null,
    created_at timestamp,
    updated_at timestamp
);
