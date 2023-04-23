CREATE TABLE IF NOT EXISTS access_count (
    id serial primary key,
    count integer not null default 0
);
