create table if not exists exercise (
     id varchar(36) primary key not null,
     name varchar(100),
     meta varchar(500),
     category varchar(100),
     type varchar(100),
     created_at timestamp not null default now(),
     updated_at timestamp not null default now()
)
