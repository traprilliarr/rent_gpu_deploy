create table if not exists "order"
(
    id              text not null
    constraint order_pk
    primary key,
    email           text not null,
    telegram        text,
    hash            text,
    value           integer,
    payment_id      text,
    payment_address text,
    ssh_key         text not null,
    gpu_fk          text
    constraint gpu_fk
    references gpu,
    user_fk         text not null
    constraint user__fk
    references "user",
    created_at      timestamp,
    updated_at      timestamp
);
