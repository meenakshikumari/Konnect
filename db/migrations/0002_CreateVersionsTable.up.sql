create table versions (
    id bigserial primary key not null,
    name varchar(50) not null,
    description varchar(500) not null,
    published boolean default false,
    service_id bigint references services(id) not null,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
);

CREATE INDEX trgm_idx_versions_name ON versions USING gin (name gin_trgm_ops);