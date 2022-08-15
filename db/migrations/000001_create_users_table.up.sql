create table if not exists users (
       id bigint primary key not null,
       first_name varchar(100),
       last_name varchar(150),
       created_at timestamp not null default now(),
       updated_at timestamp not null default now()
)