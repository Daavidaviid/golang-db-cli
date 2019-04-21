create extension if not exists "uuid-ossp";

create or replace function trigger_set_timestamp()
returns trigger as $$
begin
  new.updated_at = now();
  return new;
end;
$$ language plpgsql;

-- USER
create table users(
  id                uuid primary key default uuid_generate_v4(),
  first_name        text not null,
  last_name         text not null,
  photo_url         text,
  email             text not null unique,
  hashed_password   text not null,
  phone             text not null unique,
  deleted_at        timestamp default null,
  created_at        timestamp not null default now(),
  updated_at        timestamp not null default now()
);
create trigger set_timestamp
before update on users
for each row
execute procedure trigger_set_timestamp();

-- ARTICLE
create table articles(
  id          uuid primary key default uuid_generate_v4(),
  name        text not null,
  content        text,
  deleted_at  timestamp default null,
  created_at  timestamp not null default now(),
  updated_at  timestamp not null default now()
);
create trigger set_timestamp
before update on articles
for each row
execute procedure trigger_set_timestamp();