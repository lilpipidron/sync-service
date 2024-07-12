create table if not exists users (
    id serial primary key ,
    client_name varchar,
    version int,
    image varchar,
    cpu varchar,
    memory varchar,
    priority float8,
    need_restart bool,
    spawned_at timestamp,
    created_at timestamp,
    updated_at timestamp
);

create table if not exists AlgorithmStatus (
  id serial primary key,
  client_id bigint references users (id) on delete cascade
);