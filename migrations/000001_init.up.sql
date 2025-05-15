BEGIN;

CREATE EXTENSION IF NOT EXISTS "pgcrypto"; -- для gen_random_uuid()

CREATE TABLE tariffs
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        TEXT NOT NULL,
    description TEXT
);

CREATE TABLE tariff_resources
(
    id        UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tariff_id UUID REFERENCES tariffs (id) ON DELETE CASCADE

);

CREATE TABLE tariff_assignments
(
    id        UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id   UUID NOT NULL, -- з Account Service
    tariff_id UUID REFERENCES tariffs (id) ON DELETE CASCADE
);
COMMIT;