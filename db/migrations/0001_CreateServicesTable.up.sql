CREATE EXTENSION IF NOT EXISTS pg_trgm;

create table services (
    id bigserial primary key not null,
    name varchar(50) not null,
    description varchar(500) not null,
    published boolean default false,
    version_count INTEGER default 0,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
);

CREATE INDEX trgm_idx_services_name ON services USING gin (name gin_trgm_ops);