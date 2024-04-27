create table if not exists gpu
(
    id          text    not null
    constraint gpu_pk
    primary key,
    gpu_name    text    not null,
    price       integer not null,
    link        text,
    network     text,
    cpu         text,
    memory      text,
    storage     text,
    description text,
    available   boolean,
    created_at  timestamp,
    updated_at  timestamp
);
