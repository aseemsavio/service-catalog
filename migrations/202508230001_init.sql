-- +goose Up
CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS services
(
    service_uuid
    UUID
    PRIMARY
    KEY
    DEFAULT
    uuid_generate_v4
(
),
    name TEXT NOT NULL UNIQUE,
    description VARCHAR
(
    150
) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW
(
),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW
(
),
    CONSTRAINT description_len CHECK
(
    char_length
(
    description
) <= 150)
    );

CREATE TABLE IF NOT EXISTS versions
(
    id
    UUID
    PRIMARY
    KEY
    DEFAULT
    uuid_generate_v4
(
),
    service_uuid UUID NOT NULL REFERENCES services
(
    service_uuid
) ON DELETE CASCADE,
    name TEXT NOT NULL,
    published_on DATE NOT NULL
    );

CREATE INDEX IF NOT EXISTS idx_services_name ON services (name);
CREATE INDEX IF NOT EXISTS idx_versions_service_published ON versions (service_uuid, published_on DESC);

-- +goose Down
DROP TABLE IF EXISTS versions;
DROP TABLE IF EXISTS services;
